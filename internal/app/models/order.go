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
	OrderId        int64  `json:"orderNumber"`
	Date           string `json:"date"`
	TotalPrice     int64  `json:"totalPrice"`
	RestaurantName string `json:"restName"`
	Status         string `json:"status"`
}

type GetUserOrdersResp struct {
	Orders []ShortOrder `json:"orderList"`
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

type GetUserOrderStatusesReq struct {
	UserId int64
}

type OrderStatus struct {
	OrderId int64  `json:"id"`
	Status  string `json:"status"`
}

type GetUserOrderStatusesResp struct {
	OrderStatuses []OrderStatus `json:"statuses"`
}

// /////////////
type GetUserOrderUcaseReq struct {
	UserId  int64
	OrderId int64
}

type OrderPositionUcaseResp struct {
	Name        string
	Description string
	Count       int64
	Price       int64
	Calories    int64
	Weigth      int64
	ImagePath   string
}

type GetUserOrderUcaseResp struct {
	OrderId        int64
	Date           string
	TotalPrice     int64
	RestaurantName string
	RestaurantSlug string
	Address        string
	Status         string
	Cart           []OrderPositionUcaseResp
}

type OrderPositionResp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int64  `json:"count"`
	Price       int64  `json:"price"`
	Calories    int64  `json:"calories"`
	Weigth      int64  `json:"weight"`
	ImagePath   string `json:"imgPath"`
}
type GetUserOrderResp struct {
	OrderId        int64               `json:"orderNumber"`
	Address        string              `json:"address"`
	Date           string              `json:"date"`
	TotalPrice     int64               `json:"totalPrice"`
	RestaurantName string              `json:"restName"`
	RestaurantSlug string              `json:"restSlug"`
	Status         string              `json:"status"`
	Cart           []OrderPositionResp `json:"cart"`
}
