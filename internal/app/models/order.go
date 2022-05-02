package models

type OrderPosition struct {
	Id    int64 `json:"id"`
	Count int64 `json:"count"`
}

type OrderReq struct {
	Address string          `json:"address" valid:"address,required"`
	Comment string          `json:"comment" valid:"comment"` //TODO валидатор сделать
	Cart    []OrderPosition `json:"cart" valid:"required"`
}

type OrderUcaseReq struct {
	UserId  UserId
	Address string
	Comment string
	Cart    []OrderPosition
}

type OrderUcaseResp struct {
	OrderId int64
}

type OrderResp struct {
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

type ShortOrderResp struct {
	OrderId        int64
	Date           string
	TotalPrice     int64
	RestaurantName string
	Status         string
}

type ShortOrder struct {
	OrderId        int64  `json:"id"`
	Date           string `json:"date"`
	TotalPrice     int64  `json:"totalPrice"`
	RestaurantName string `json:"restName"`
	Status         string `json:"status"`
}

type GetUserOrdersResp struct {
	Orders []ShortOrder `json:"orders"`
}

////////////

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

type GetUserOrderStatusesReq struct {
	UserId int64
}

type OrderStatus struct {
	OrderId int64
	Status  string
}

type GetUserOrderStatusesResp struct {
	OrderStatuses []OrderStatus `json:"orders"`
}
