package repository

import (
	"database/sql"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	DB *sqlx.DB
}

const (
	ordersMax = 100
)

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func makeAddOrderQuery(order *models.CreateOrderRepoReq) string {
	if len(order.Cart) == 0 {
		return `INSERT INTO orders_internal (user_id , address, comment, cart) VALUES ($1,$2,$3)`
	}
	query := `INSERT INTO orders_internal (user_id , address, comment, promocode_id, cart) VALUES ($1,$2,$3, (SELECT id FROM promocodes WHERE promocode=$4), ARRAY[`

	nextPlaceholderNum := 5
	for i := 0; i < len(order.Cart)-1; i++ {
		query += `($` + strconv.Itoa(nextPlaceholderNum)
		nextPlaceholderNum++
		query += `,$` + strconv.Itoa(nextPlaceholderNum)
		nextPlaceholderNum++
		query += `,$` + strconv.Itoa(nextPlaceholderNum) + `)::order_dish,`
		nextPlaceholderNum++
	}
	query += `($` + strconv.Itoa(nextPlaceholderNum)
	nextPlaceholderNum++
	query += `,$` + strconv.Itoa(nextPlaceholderNum)
	nextPlaceholderNum++
	query += `,$` + strconv.Itoa(nextPlaceholderNum) + `)::order_dish]) RETURNING id`
	return query
}

func expandOrder(order *models.CreateOrderRepoReq) []interface{} {
	args := make([]interface{}, 0, len(order.Cart)*3+4)
	args = append(args, order.UserId, order.Address, order.Comment, order.Promocode)
	for _, orderPos := range order.Cart {
		args = append(args, orderPos.Id, orderPos.Count, 0)
	}
	return args
}

func (r *OrderRepo) CreateOrder(order *models.CreateOrderRepoReq) (*models.CreateOrderRepoResp, error) {
	query := makeAddOrderQuery(order)

	var newOrderId int64
	err := r.DB.QueryRow(query, expandOrder(order)...).Scan(&newOrderId)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	if newOrderId == 0 {
		return nil, servErrors.NewError(servErrors.DB_INSERT, "")
	}
	return &models.CreateOrderRepoResp{OrderId: newOrderId}, nil
}

func (r *OrderRepo) GetUserOrders(user *models.GetUserOrdersRepoReq) (*models.GetUserOrdersRepoResp, error) {
	orders := make([]*models.ShortOrderRepo, 0, ordersMax)
	err := r.DB.Select(&orders, `SELECT id, date, restaurant_name, total_price_discount, status, delivery_price FROM orders WHERE user_id = $1 ORDER BY fulldate DESC `, user.UserId)

	ordersResp := make([]models.ShortOrderRepo, len(orders))
	for i, order := range orders {
		ordersResp[i] = *order
	}
	switch err {
	case nil:
		return &models.GetUserOrdersRepoResp{Orders: ordersResp}, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *OrderRepo) GetUserOrderStatuses(user *models.GetUserOrderStatusesRepoReq) (*models.GetUserOrderStatusesRepoResp, error) {
	orderStatuses := make([]*models.OrderStatusRepo, 0, ordersMax)
	err := r.DB.Select(&orderStatuses, `SELECT id, status FROM orders WHERE user_id = $1 ORDER BY fulldate DESC`, user.UserId)

	ordersResp := make([]models.OrderStatusRepo, len(orderStatuses))
	for i, order := range orderStatuses {
		ordersResp[i] = *order
	}
	switch err {
	case nil:
		return &models.GetUserOrderStatusesRepoResp{OrderStatuses: ordersResp}, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *OrderRepo) GetUserOrder(req *models.GetUserOrderRepoReq) (*models.GetUserOrderRepoResp, error) {
	order := &models.GetUserOrderRepoResp{}
	err := r.DB.Get(order, `SELECT id, address, user_id, date, restaurant_name, restaurant_slug, total_price, total_price_discount, delivery_price, status FROM orders WHERE id = $1 `, req.OrderId)
	cart := make([]*models.OrderPositionRepoResp, 0)
	if err == nil {
		err = r.DB.Select(&cart, `SELECT d.description, d.name, d.weight, d.calories, d.image_path, c.price price, c.count from dishes d JOIN (SELECT (unnest(cart)::order_dish).id as id, (unnest(cart)::order_dish).count, (unnest(cart)::order_dish).price FROM orders_internal WHERE id=$1) c ON d.id=c.id`, req.OrderId)
		if err == nil {
			order.Cart = make([]models.OrderPositionRepoResp, len(cart))
			for i, poz := range cart {
				order.Cart[i] = *poz
			}
		}
	}
	switch err {
	case nil:
		return order, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *OrderRepo) GetAddress(req *models.GetAddressRepoReq) (*models.GetAddressRepoResp, error) {
	city := &struct {
		Id   int64
		Name string
	}{}
	street := &struct {
		Id   int64
		Name string
	}{}
	house := &struct {
		House string
	}{}
	err := r.DB.Get(city, `SELECT id, name FROM cities WHERE name ILIKE $1`, req.City)
	if err == nil {
		err = r.DB.Get(street, `SELECT id, name FROM streets WHERE city_id = $1 AND main_name ILIKE $2 AND type ILIKE $3`, city.Id, req.Street, req.StreetType)
	}
	if err == nil {
		err = r.DB.Get(house, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2`, street.Id, req.House)
	}

	switch err {
	case nil:
		return &models.GetAddressRepoResp{
			City:   city.Name,
			Street: street.Name,
			House:  house.House,
		}, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ADDRESS, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
