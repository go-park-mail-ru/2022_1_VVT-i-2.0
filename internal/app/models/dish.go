package models

type Dish struct {
	Id           int
	RestaurantId string
	Name         string
	Description  string
	Image_path   string
	Calories     int
	Price        int
}

type DishDataStorage struct {
	Id           int
	RestaurantId int `db:"restaurant_id"`
	Name         string
	Description  string
	ImagePath    string `db:"image_path"`
	Calories     int
	Price        int
	Weight       int
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

type DishesDataStorage struct {
	Dishes []DishDataStorage
}

type DishesUcase struct {
	Dishes []DishUcase
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

type RestaurantDishesResp struct {
	Id            	int        	`json:"id"`
	Name          	string     	`json:"restName"`
	ImagePath     	string     	`json:"imgPath"`
	Slug          	string     	`json:"slug"`
	MinPrice      	int        	`json:"minPrice"`
	Rating        	float64    	`json:"rating"`
	TimeToDelivery 	string     	`json:"timeToDeliver"`
	ReviewCount		int			`json:"reviewCount"`
	Dishes         	[]DishResp 	`json:"dishes"`
}
