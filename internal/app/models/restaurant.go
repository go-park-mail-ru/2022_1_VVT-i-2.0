package models

//// Repository

// Repository request

type GetRestaurantByCategoryRepoReq struct {
	Name string
}

type GetRestaurantBySearchQueryRepoReq struct {
	Query string
}

// Repository models

type RestaurantRepo struct {
	Id                   int `db:"id"`
	Name                 string
	ImagePath            string `db:"image_path"`
	Slug                 string
	MinPrice             int `db:"min_price"`
	AggRating            int `db:"agg_rating"`
	ReviewCount          int `db:"review_count"`
	UpMinutsToDelivery   int `db:"up_time_to_delivery"`
	DownMinutsToDelivery int `db:"down_time_to_delivery"`
}

type RestaurantsRepo struct {
	Restaurants []RestaurantRepo
}

//// UseCase

// UseCase request

type GetRestaurantByCategoryUcaseReq struct {
	Name string
}

type GetRestaurantBySearchQueryUcaseReq struct {
	Query string
}

// UseCase models

type RestaurantUcase struct {
	Id                   int
	Name                 string
	ImagePath            string
	Slug                 string
	MinPrice             int
	AggRating            int
	ReviewCount          int
	UpMinutsToDelivery   int
	DownMinutsToDelivery int
}

type RestaurantsUcase struct {
	Restaurants []RestaurantUcase
}

//// Handler

//easyjson:json
type RestaurantResp struct {
	Id             int     `json:"id"`
	Name           string  `json:"restName"`
	ImagePath      string  `json:"imgPath"`
	Slug           string  `json:"slug"`
	MinPrice       int     `json:"price"`
	Rating         float64 `json:"rating"`
	TimeToDelivery string  `json:"timeToDeliver"`
}

//easyjson:json
type AllRestaurantsResp struct {
	Restaurants []RestaurantResp `json:"restaurants"`
}
