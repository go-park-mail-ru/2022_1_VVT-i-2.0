package models

import (
	"database/sql"
)

type OrderPositionStorage struct {
	Id    int64
	Count int64
	Price int64
}

type OrderStorage struct {
	OrderId    int64
	UserId     int64
	Address    string
	Comment    sql.NullString
	TotalPrice string
	Cart       []OrderPositionStorage
}

type OrderPositionUcase struct {
	Id    int64
	Count int64
}

type OrderPositionRepo struct {
	Id    int64
	Count int64
}

type CreateOrderRepoReq struct {
	UserId  int64
	Address string
	Comment string
	Cart    []OrderPositionRepo
}

type CreateOrderRepoResp struct {
	OrderId int64
}

type CreateOrderUcaseReq struct {
	UserId  int64
	Address string
	Comment string
	Cart    []OrderPositionUcase
}

type CreateOrderUcaseResp struct {
	OrderId int64
}

type GetUserOrdersUcaseReq struct {
	UserId int64
}

type ShortOrderUcase struct {
	OrderId        int64
	Date           string
	TotalPrice     int64
	RestaurantName string
	Status         string
}

type GetUserOrdersUcaseResp struct {
	Orders []ShortOrderUcase
}

type GetUserOrdersRepoReq struct {
	UserId int64
}

type ShortOrderRepo struct {
	OrderId        int64 `db:"id"`
	Date           string
	TotalPrice     int64  `db:"total_price"`
	RestaurantName string `db:"restaurant_name"`
	Status         string
}

type GetUserOrdersRepoResp struct {
	OrderStatuses []ShortOrderRepo
}

type GetUserOrderStatusesUcaseReq struct {
	UserId int64
}

type OrderStatusUcase struct {
	OrderId int64
	Status  string
}

type GetUserOrderStatusesUcaseResp struct {
	OrderStatuses []OrderStatusUcase
}

type GetUserOrderStatusesRepoReq struct {
	UserId int64
}

type OrderStatusRepo struct {
	OrderId int64 `db:"id"`
	Status  string
}

type GetUserOrderStatusesRepoResp struct {
	OrderStatuses []OrderStatusRepo
}
