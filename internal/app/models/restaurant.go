package models

type Restaurant struct {
	Id         int
	Name       string
	City       string
	Address    string
	Image_path string
	Slug       string
	Min_price  int
	Avg_price  int
	Rating     int
	Count_rating int
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

type RestaurantsDataStorage struct {
	Restaurants []RestaurantDataStorage
}

type RestaurantUsecase struct {
	Id         int
	Name       string
	City       string
	Address    string
	Image_path string
	Slug       string
	Min_price  int
	Avg_price  int
	Rating     int
	Count_rating int
}

type RestaurantsUsecase struct {
	Restaurants []RestaurantUsecase
}

// imgPath, restName, slug, timeToDeliver, price, rating
type RestaurantJson struct {
	Id            int     `json:"id"`
	Name          string  `json:"restName"`
	City          string  `json:"city"`
	Address       string  `json:"address"`
	Image_path    string  `json:"imgPath"`
	Slug          string  `json:"slug"`
	Min_price     int     `json:"minPrice"`
	Avg_price     int     `json:"avgPrice"`
	Rating        float64 `json:"rating"`
	TimeToDeliver string  `json:"timeToDeliver"`
}

type RestaurantsResponse struct {
	Restaurants []RestaurantJson `json:"restaurants"`
}

type RestaurantsDishJson struct {
	Id         int        `json:"id"`
	Name       string     `json:"restName"`
	City       string     `json:"city"`
	Address    string     `json:"address"`
	Image_path string     `json:"imgPath"`
	Slug       string     `json:"slug"`
	Min_price  int        `json:"minPrice"`
	Avg_price  int        `json:"avgPrice"`
	Dishes     []DishJson `json:"dishes"`
}

type RestaurantsDishesJsonForKirill struct {
	Id 			int			`json:"id"`
	Name        string		`json:"restName"`
	City 		string		`json:"city"`
	Address 	string		`json:"address"`
	Image_path 	string		`json:"imgPath"`
	Slug 		string		`json:"slug"`
	Min_price 	int			`json:"minPrice"`
	Avg_price 	int			`json:"avgPrice"`
	Rating 		float64		`json:"rating"`
	TimeToDelivery string 	`json:"timeToDeliver"`
	Dishes		[]DishJsonForKirill 	`json:"dishes"`
}

type RestaurantJsonForKirill struct {
	Id 				int		`json:"id"`
	Name        	string	`json:"restName"`
	City 			string	`json:"city"`
	Address 		string	`json:"address"`
	Image_path 		string	`json:"imgPath"`
	Slug 			string	`json:"slug"`
	Min_price 		int		`json:"min_price"`
	Avg_price 		int		`json:"price"`
	Rating 			float64	`json:"rating"`
	TimeToDelivery 	string `json:"timeToDeliver"`
}

type RestaurantsResponseForKirill struct {
	Restaurants []RestaurantJsonForKirill 	`json:"restaurants"`
}