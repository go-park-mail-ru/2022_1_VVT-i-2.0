package models

type RestaurantUsecase struct {
	Id            int     `json:"id"`
	Name          string  `json:"restName"`
	City          string  `json:"city"`
	Address       string  `json:"address"`
	Image_path    string  `json:"imgPath"`
	Slug          string  `json:"slug"`
	Min_price     int     `json:"minPrice"`
	Avg_price     int     `json:"avgPrice"`
	Rating        int 	  `json:"rating"`
	TimeToDeliver string  `json:"timeToDeliver"`
}

type RestaurantsUsecase struct {
	Restaurants []RestaurantUsecase
}

type DishUseCase struct {
	Id 			int		`json:"id"`
	Restaurant 	int		`json:"restaurany"`
	Name 		string	`json:"productName"`
	Description string	`json:"description"`
	Image_path 	string	`json:"imgPath"`
	Calories 	int		`json:"info"`
	Price 		int		`json:"price"`
	Weight		int		`json:"weight"`
}

type DishesUseCase struct {
	Dishes []DishUseCase
}
