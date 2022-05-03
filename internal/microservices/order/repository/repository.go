package repository

import (
	"database/sql"
	"fmt"
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
	// INSERT INTO orders (user_id , address, comment, cart) VALUES ($1,$2,$3, ARRAY[($4,$5)::order_position, ($6,$7)::order_position])`
	query := `INSERT INTO orders_internal (user_id , address, comment, cart) VALUES ($1,$2,$3, ARRAY[`

	nextPlaceholderNum := 4
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

func ExpandOrder(order *models.CreateOrderRepoReq) []interface{} {
	var args []interface{}
	args = append(args, order.UserId, order.Address, order.Comment)
	for _, orderPos := range order.Cart {
		args = append(args, orderPos.Id, orderPos.Count, 0)
	}
	return args
}

func (r *OrderRepo) CreateOrder(order *models.CreateOrderRepoReq) (*models.CreateOrderRepoResp, error) {
	query := makeAddOrderQuery(order)
	fmt.Println(query)
	// вычислить заначение стоимости

	var newOrderId int64
	err := r.DB.QueryRow(query, ExpandOrder(order)...).Scan(&newOrderId)
	fmt.Println(err)

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
	err := r.DB.Select(&orders, `SELECT id, date, restaurant_name, total_price, status FROM orders WHERE user_id = $1 ORDER BY fulldate DESC `, user.UserId)

	ordersResp := make([]models.ShortOrderRepo, len(orders))
	for i, order := range orders {
		ordersResp[i] = *order
	}
	switch err {
	case nil:
		return &models.GetUserOrdersRepoResp{OrderStatuses: ordersResp}, nil
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
	// TODO: hardcode
	order := &models.GetUserOrderRepoResp{}
	// забрать данные заказа
	// пройтись по массиву товаров в корзине и взять данные товаров
	fmt.Println(`SELECT id, user_id, date, restaurant_name, total_price, status, cart FROM orders WHERE id = `, req.OrderId)
	err := r.DB.Get(order, `SELECT id, address, user_id, date, restaurant_name, total_price, status FROM orders WHERE id = $1 `, req.OrderId)
	cart := make([]*models.OrderPositionRepoResp, 0, 100)
	if err == nil {
		fmt.Printf("SELECT d.id id, d.description, d.name, d.weight, d.calories, d.image_path, c.price price, c.count from dishes d JOIN (SELECT (unnest(cart)::order_dish).id as id, (unnest(cart)::order_dish).count, (unnest(cart)::order_dish).price FROM orders_internal WHERE id=%v) c ON d.id=c.id\n", req.OrderId)
		err = r.DB.Select(&cart, `SELECT d.description, d.name, d.weight, d.calories, d.image_path, c.price price, c.count from dishes d JOIN (SELECT (unnest(cart)::order_dish).id as id, (unnest(cart)::order_dish).count, (unnest(cart)::order_dish).price FROM orders_internal WHERE id=$1) c ON d.id=c.id`, req.OrderId)
		fmt.Println(cart[0])
		fmt.Println(cart[1])
		if err == nil {
			order.Cart = make([]models.OrderPositionRepoResp, len(cart))
			for i, poz := range cart {
				order.Cart[i] = *poz
			}
		}
	}
	fmt.Println("----------")
	fmt.Println(order)
	fmt.Println(err)
	switch err {
	case nil:
		// return &models.GetUserOrdersRepoResp{OrderStatuses: ordersResp}, nil
		return order, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
