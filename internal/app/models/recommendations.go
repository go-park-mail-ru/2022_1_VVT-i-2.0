package models

//easyjson:json
type RecommendationsOrderPosition struct {
	Id int64 `json:"id"`
}

//easyjson:json
type RecommendationsReq struct {
	RestId    int64                          `json:"restId"`
	OrderList []RecommendationsOrderPosition `json:"orderList"`
}

//easyjson:json
type RecommendationResp struct {
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

type RecommendationsResp struct {
	Dishes []RecommendationResp `json:"dishes"`
}

type RecommendationsUcaseReq struct {
	Limit    int
	RestId   int64
	DishesId []int64
}

type RecommendationsRepoReq struct {
	Limit    int
	RestId   int64
	DishesId []int64
}

type RecommendationRepo struct {
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

type RecommendationsRepoResp struct {
	Dishes []RecommendationRepo
}

type RecommendationUcase struct {
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

type RecommendationsUcaseResp struct {
	Dishes []RecommendationUcase
}
