package models

//easyjson:json
type OrderPosition struct {
	Id    int64 `json:"id"`
	Count int64 `json:"count"`
}

//easyjson:json
type OrderReq struct {
	Address   string          `json:"address" valid:"address,required"`
	Entrance  string          `json:"entrance" valid:"utfletternum"`
	Intercom  string          `json:"intercom" valid:"utfletternum"`
	Floor     string          `json:"floor" valid:"utfletternum"`
	Flat      string          `json:"flat" valid:"utfletternum"`
	Comment   string          `json:"comment" valid:"comment"`
	Cart      []OrderPosition `json:"cart" valid:"required"`
	Promocode string          `json:"promocode" valid:"promocode"`
}

type OrderUcaseReq struct {
	UserId    int64
	Address   string
	Entrance  string
	Intercom  string
	Floor     string
	Flat      string
	Comment   string
	Promocode string
	Cart      []OrderPosition
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

//easyjson:json
type ShortOrder struct {
	OrderId        int64  `json:"orderNumber"`
	Date           string `json:"date"`
	TotalPrice     int64  `json:"totalPrice"`
	RestaurantName string `json:"restName"`
	Status         string `json:"status"`
}

//easyjson:json
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

//easyjson:json
type OrderStatus struct {
	OrderId int64  `json:"id"`
	Status  string `json:"status"`
}

//easyjson:json
type GetUserOrderStatusesResp struct {
	OrderStatuses []OrderStatus `json:"statuses"`
}

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
	Discount       int64
	DeliveryPrice  int64
	RestaurantName string
	RestaurantSlug string
	Address        string
	Status         string
	Cart           []OrderPositionUcaseResp
}

//easyjson:json
type OrderPositionResp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int64  `json:"count"`
	Price       int64  `json:"price"`
	Calories    int64  `json:"calories"`
	Weigth      int64  `json:"weight"`
	ImagePath   string `json:"imgPath"`
}

//easyjson:json
type GetUserOrderResp struct {
	OrderId        int64               `json:"orderNumber"`
	Address        string              `json:"address"`
	Date           string              `json:"date"`
	TotalPrice     int64               `json:"totalPrice"`
	Discount       int64               `json:"summaryDiscount"`
	DeliveryPrice  int64               `json:"deliveryPrice"`
	RestaurantName string              `json:"restName"`
	RestaurantSlug string              `json:"restSlug"`
	Status         string              `json:"status"`
	Cart           []OrderPositionResp `json:"cart"`
}
