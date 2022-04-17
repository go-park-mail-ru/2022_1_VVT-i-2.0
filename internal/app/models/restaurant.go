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
	Rating     float64
}

type RestaurantDataStorage struct {
	Id         int
	Name       string
	City       string
	Address    string
	Image_path string
	Slug       string
	Min_price  int
	Avg_price  int
	Rating     float64
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
	Rating     float64
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

type Dish struct {
	Id          int
	Restaurant  string
	Name        string
	Description string
	Image_path  string
	Calories    int
	Price       int
}

type DishDataStorage struct {
	Id          int
	Restaurant  string
	Name        string
	Description string
	Image_path  string
	Calories    int
	Price       int
}

type DishesDataStorage struct {
	Dishes []DishDataStorage
}

// TODO: они ждут info!!
type DishJson struct {
	Id          int    `json:"id"`
	Restaurant  string `json:"restaurant"`
	Name        string `json:"productName"`
	Description string `json:"description"`
	Image_path  string `json:"imgPath"`
	Calories    int    `json:"calories"`
	Price       int    `json:"price"`
	Info        string `json:"info"`
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
