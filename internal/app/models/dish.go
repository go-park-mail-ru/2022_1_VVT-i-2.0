package models

type GetRestaurantDishesRepoReq struct {
	Id int
}

type DishRepo struct {
	Id           int
	RestaurantId int `db:"restaurant_id"`
	Name         string
	Description  string
	ImagePath    string `db:"image_path"`
	Calories     int
	Price        int
	Weight       int
}

type DishCategoriRepo struct {
	Id           int
	Categori		int `db:"categori"`
	RestaurantId int `db:"restaurant_id"`
	Name         string
	Description  string
	ImagePath    string `db:"image_path"`
	Calories     int
	Price        int
	Weight       int
}

type GetRestaurantDishesCategoriesRepoResp struct {
	Dishes []DishCategoriRepo
}

type GetRestaurantDishesRepoResp struct {
	Dishes []DishRepo
}

type GetRestaurantDishesUcaseReq struct {
	Slug string
}

type DishUcase struct {
	Id           int
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
	Dishes     []DishCategoriRepo
}

type GetRestaurantDishesCategoriesUcaseResp struct {
	Id                   int
	Name                 string
	ImagePath            string
	Slug                 string
	MinPrice             int
	AggRating            int
	ReviewCount          int
	UpMinutsToDelivery   int
	DownMinutsToDelivery int
	Dishes               []CategoriesDishesUcaseResp
}

type GetRestaurantDishesUcaseResp struct {
	Id                   int
	Name                 string
	ImagePath            string
	Slug                 string
	MinPrice             int
	AggRating            int
	ReviewCount          int
	UpMinutsToDelivery   int
	DownMinutsToDelivery int
	Dishes               []DishUcase
}

type DishResp struct {
	Id           int    `json:"id"`
	RestaurantId int    `json:"restaurant"`
	Name         string `json:"productName"`
	Description  string `json:"description"`
	ImagePath    string `json:"imgPath"`
	Calories     int    `json:"info"`
	Price        int    `json:"price"`
	Weight       int    `json:"weight"`
}

type GetRestaurantDishesResp struct {
	Id             int        `json:"id"`
	Name           string     `json:"restName"`
	ImagePath      string     `json:"imgPath"`
	Slug           string     `json:"slug"`
	MinPrice       int        `json:"minPrice"`
	Rating         float64    `json:"rating"`
	TimeToDelivery string     `json:"timeToDeliver"`
	ReviewCount    int        `json:"reviewCount"`
	Dishes         []DishResp `json:"dishes"`
}

type CategoriesDishesDelivery struct {
	Categories string
	Dishes     []DishCategoriRepo
}

type GetRestaurantDishesCategoriesResp struct {
	Id             int        `json:"id"`
	Name           string     `json:"restName"`
	ImagePath      string     `json:"imgPath"`
	Slug           string     `json:"slug"`
	MinPrice       int        `json:"minPrice"`
	Rating         float64    `json:"rating"`
	TimeToDelivery string     `json:"timeToDeliver"`
	ReviewCount    int        `json:"reviewCount"`
	Dishes         []CategoriesDishesDelivery `json:"dishes"`
}
