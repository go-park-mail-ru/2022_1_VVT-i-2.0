package models

type Id uint64

type GetRestaurantBySlugRepoReq struct {
	Slug string
}

type GetCategoriesByIdRepoReq struct {
	Id Id
}

type GetRestaurantDishesRepoReq struct {
	Id Id
}

type DishesRestaurantRepo struct {
	Id                    int `db:"id"`
	Name                  string
	ImagePath             string `db:"image_path"`
	Slug                  string
	MinPrice              int `db:"min_price"`
	AggRating             int `db:"agg_rating"`
	ReviewCount           int `db:"review_count"`
	UpMinutesToDelivery   int `db:"up_time_to_delivery"`
	DownMinutesToDelivery int `db:"down_time_to_delivery"`
}

type Categories struct {
	Categories []string `db:"categories"`
}

type DishCategoriesRepo struct {
	Id           int
	Category     int `db:"category"`
	RestaurantId int `db:"restaurant_id"`
	Name         string
	Description  string
	ImagePath    string `db:"image_path"`
	Calories     int
	Price        int
	Weight       int
}

type GetRestaurantDishesCategoriesRepoResp struct {
	Dishes []DishCategoriesRepo
}

type GetRestaurantDishesUcaseReq struct {
	Slug string
}

type DishCategoriesUsecase struct {
	Id           int
	Category     int
	RestaurantId int
	Name         string
	Description  string
	ImagePath    string
	Calories     int
	Price        int
	Weight       int
}

type CategoriesDishesUcaseResp struct {
	Categories string
	Dishes     []int
}

type GetRestaurantDishesCategoriesUcaseResp struct {
	Id                    int
	Name                  string
	ImagePath             string
	Slug                  string
	MinPrice              int
	AggRating             int
	ReviewCount           int
	UpMinutesToDelivery   int
	DownMinutesToDelivery int
	Dishes                []DishCategoriesUsecase
	Categories            []CategoriesDishesUcaseResp
}

//easyjson:json
type DishCategoriesResp struct {
	Id           int    `json:"id"`
	Category     int    `json:"category"`
	RestaurantId int    `json:"restaurant"`
	Name         string `json:"productName"`
	Description  string `json:"description"`
	ImagePath    string `json:"imgPath"`
	Calories     int    `json:"info"`
	Price        int    `json:"price"`
	Weight       int    `json:"weight"`
}

//easyjson:json
type CategoriesDishesDelivery struct {
	Category string `json:"category"`
	Dishes   []int  `json:"dishes"`
}

//easyjson:json
type GetRestaurantDishesCategoriesResp struct {
	Id             int                        `json:"id"`
	Name           string                     `json:"restName"`
	ImagePath      string                     `json:"imgPath"`
	Slug           string                     `json:"slug"`
	MinPrice       int                        `json:"minPrice"`
	Rating         float64                    `json:"rating"`
	TimeToDelivery string                     `json:"timeToDeliver"`
	ReviewCount    int                        `json:"reviewCount"`
	Dishes         []DishCategoriesResp       `json:"dishes"`
	Categories     []CategoriesDishesDelivery `json:"categories"`
}
