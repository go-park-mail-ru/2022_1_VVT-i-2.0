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
<<<<<<< HEAD
	Id         int
	Name       string
	City       string
	Address    string
	Image_path string
	Slug       string
	Min_price  int
	Avg_price  int
	Rating     float64
=======
	Id   		int
	Name 		string
	City 		string
	Address 	string
	Image_path 	string
	Slug 		string
	Min_price 	int
	Avg_price 	int
	Rating 		float64
>>>>>>> 56c01c4385118e330bd2e349e9814480f8d0e6e6
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
	Id 			int
	Restaurant 	int
	Name 		string
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
<<<<<<< HEAD
	Id          int    `json:"id"`
	Restaurant  string `json:"restaurant"`
	Name        string `json:"productName"`
	Description string `json:"description"`
	Image_path  string `json:"imgPath"`
	Calories    int    `json:"calories"`
	Price       int    `json:"price"`
	Info        string `json:"info"`
=======
	Id 			int		`json:"id"`
	Restaurant 	int	`json:"restaurany"`
	Name 		string	`json:"name"`
	Description string	`json:"description"`
	Image_path 	string	`json:"image_path"`
	Calories 	int		`json:"calories"`
	Price 		int		`json:"price"`
>>>>>>> 56c01c4385118e330bd2e349e9814480f8d0e6e6
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

type DishJsonForKirill struct {
	Id 			int		`json:"id"`
	Restaurant 	int	`json:"restaurany"`
	Name 		string	`json:"productName"`
	Description string	`json:"description"`
	Image_path 	string	`json:"imgPath"`
	Calories 	int		`json:"info"`
	Price 		int		`json:"price"`
}

type RestaurantsDishJsonForKirill struct {
	Id 			int			`json:"id"`
	Name        string		`json:"restName"`
	City 		string		`json:"city"`
	Address 	string		`json:"address"`
	Image_path 	string		`json:"imgPath"`
	Slug 		string		`json:"slug"`
	Min_price 	int			`json:"min_price"`
	Avg_price 	int			`json:"price"`
	Rating 		float64		`json:"rating"`
	TimeToDelivery string `json:"timeToDeliver"`
	Dishes		[]DishJsonForKirill 	`json:"dishes"`
}

type RestaurantJsonForKirill struct {
	Id 			int		`json:"id"`
	Name        string	`json:"restName"`
	City 		string	`json:"city"`
	Address 	string	`json:"address"`
	Image_path 	string	`json:"imgPath"`
	Slug 		string	`json:"slug"`
	Min_price 	int		`json:"min_price"`
	Avg_price 	int		`json:"price"`
	Rating 		float64	`json:"rating"`
	TimeToDelivery string `json:"timeToDeliver"`
}

type RestaurantsResponseForKirill struct {
	Restaurants []RestaurantJsonForKirill 	`json:"restaurants"`
}

type CommentRestaurantDataStorage struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
}

type AddCommentRestaurantDataStorage struct {
	Restaurant	int
	User_id		int
	Comment_text string
}
