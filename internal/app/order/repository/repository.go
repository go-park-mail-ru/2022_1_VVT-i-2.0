package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	DB *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func makeAddOrderQuery(order *models.OrderRepoInput) string {
	// INSERT INTO orders (user_id , address, comment, cart) VALUES ($1,$2,$3, ARRAY[($4,$5)::order_position, ($6,$7)::order_position])`
	query := `INSERT INTO orders (user_id , address, comment, cart) VALUES ($1,$2,$3, ARRAY[`

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

func ExpandOrder(order *models.OrderRepoInput) []interface{} {
	var args []interface{}
	args = append(args, order.User_id, order.Address, order.Comment)
	for _, orderPos := range order.Cart {
		args = append(args, orderPos.Id, orderPos.Count, 0)
	}
	return args
}

func (r *OrderRepo) AddOrder(order *models.OrderRepoInput) (*models.OrderRepoAnsw, error) {
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
	return &models.OrderRepoAnsw{OrderId: newOrderId}, nil
}
