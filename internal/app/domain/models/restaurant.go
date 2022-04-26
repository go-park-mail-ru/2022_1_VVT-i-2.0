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

type RestaurantDataStorage struct {
	Id   		int
	Name 		string
	City 		string
	Address 	string
	Image_path 	string
	Slug 		string
	Min_price 	int
	Avg_price 	int
	Rating     int
	Count_rating int
}

var Rest = &RestaurantDataStorage{
	Id: 1,
	Name: "Name",
	City: "City",
	Address: "Address",
	Image_path: "imgPath",
	Slug: "slug",
	Min_price: 1,
	Avg_price: 1,
	Rating: 1,
	Count_rating: 1,
}

type RestaurantsDataStorage struct {
	Restaurants []RestaurantDataStorage
}

type DishDataStorage struct {
	Id 			int
	Restaurant 	int
	Name 		string
	Description string
	Image_path  string
	Calories    int
	Price       int
	Weight		int
}

var Dish = &DishDataStorage{
	Id: 1,
	Restaurant: 1,
	Name: "Name",
	Description: "Description",
	Image_path: "imgPath",
	Calories: 1,
	Price: 1,
	Weight: 1,
}

type DishesDataStorage struct {
	Dishes []DishDataStorage
}
