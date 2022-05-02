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
