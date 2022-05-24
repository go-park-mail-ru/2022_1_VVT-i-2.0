package models

type OrderPositionUcase struct {
	Id    int64
	Count int64
}

type OrderPositionRepo struct {
	Id    int64
	Count int64
}

type CreateOrderRepoReq struct {
	UserId    int64
	Address   string
	Comment   string
	Promocode string
	Cart      []OrderPositionRepo
}

type CreateOrderRepoResp struct {
	OrderId int64
}

type CreateOrderUcaseReq struct {
	UserId    int64
	Address   string
	Comment   string
	Promocode string
	Cart      []OrderPositionUcase
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
	TotalPrice     int64  `db:"total_price_discount"`
	RestaurantName string `db:"restaurant_name"`
	Status         string
}

type GetUserOrdersRepoResp struct {
	Orders []ShortOrderRepo
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

type GetUserOrderRepoReq struct {
	OrderId int64
}

type OrderPositionRepoResp struct {
	Name        string
	Description string
	Count       int64
	Price       int64
	Calories    int64
	Weight      int64
	ImagePath   string `db:"image_path"`
}

type GetUserOrderRepoResp struct {
	UserId             int64 `db:"user_id"`
	OrderId            int64 `db:"id"`
	Date               string
	TotalPrice         int64  `db:"total_price"`
	TotalPriceDiscount int64  `db:"total_price_discount"`
	RestaurantName     string `db:"restaurant_name"`
	RestaurantSlug     string `db:"restaurant_slug"`
	Address            string
	Status             string
	Cart               []OrderPositionRepoResp
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
	RestaurantName string
	RestaurantSlug string
	Address        string
	Status         string
	Cart           []OrderPositionUcaseResp
}

type GetAddressRepoReq struct {
	City       string
	Street     string
	StreetType string
	House      string
}

type GetAddressRepoResp struct {
	City   string
	Street string
	House  string
}
