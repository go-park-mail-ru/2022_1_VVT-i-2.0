package models

//// Delivery

// Delivery Request

type RecommendationsOrderPosition struct {
	Id    int64 `json:"id"`
	Count int64 `json:"count"`
}

type RecommendationsOrderLists struct {
	RestId int64 `json:"restId"`
	OrderList []RecommendationsOrderPosition `json:"orderList"`
}

// models

type DishRecommendationDelivery struct {
	Id           	int 	`json:"id"`
	Category     	int 	`json:"category"`
	RestaurantId 	int 	`json:"restaurant"`
	Name        	string	`json:"productName"`
	Description 	string	`json:"description"`
	ImagePath   	string	`json:"imgPath"`
	Calories    	int   	`json:"info"`
	Price       	int    	`json:"price"`
	Weight      	int    	`json:"weight"`
}

type DishRecommendationListsDelivery struct {
	Dishes []DishRecommendationDelivery	`json:"dishes"`
}

//// UseCase

// UseCase Request

type RecommendationsOrderListsUsecaseReq struct {
	RestId int64
	DishesId []int64
}

// UseCase models

type DishRecommendationUsecase struct {
	Id           	int
	Category     	int
	RestaurantId 	int
	Name        	string
	Description 	string
	ImagePath   	string
	Calories    	int
	Price       	int
	Weight      	int
}

type DishRecommendationListsUsecase struct {
	Dishes []DishRecommendationUsecase
}








