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
	Id   		int
	Name 		string
	City 		string
	Address 	string
	Image_path 	string
	Slug 		string
	Min_price 	int
	Avg_price 	int
	Rating 		float64
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
	Weight		int
}

type DishUseCase struct {
	Id 			int
	Restaurant 	int
	Name 		string
	Description string
	Image_path  string
	Calories    int
	Price       int
	Weight		int
}

type DishesDataStorage struct {
	Dishes []DishDataStorage
}

type DishesUseCase struct {
	Dishes []DishUseCase
}

type DishJson struct {
	Id          int    	`json:"id"`
	Restaurant  string 	`json:"restaurant"`
	Name        string 	`json:"productName"`
	Description string 	`json:"description"`
	Image_path  string 	`json:"imgPath"`
	Calories    int    	`json:"calories"`
	Price       int    	`json:"price"`
	Weight		int		`json:"weight"`
	Info        string 	`json:"info"`
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
	Restaurant 	int		`json:"restaurany"`
	Name 		string	`json:"productName"`
	Description string	`json:"description"`
	Image_path 	string	`json:"imgPath"`
	Calories 	int		`json:"info"`
	Price 		int		`json:"price"`
	Weight		int		`json:"weight"`
}

type RestaurantsDishesJsonForKirill struct {
	Id 			int			`json:"id"`
	Name        string		`json:"restName"`
	City 		string		`json:"city"`
	Address 	string		`json:"address"`
	Image_path 	string		`json:"imgPath"`
	Slug 		string		`json:"slug"`
	Min_price 	int			`json:"min_price"`
	Avg_price 	int			`json:"price"`
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

type CommentRestaurantDataStorage struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type AddCommentRestaurantDataStorage struct {
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type AddCommentRestaurantUseCase struct {
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type AddCommentsRestaurantDataStorage struct {
	Comment []AddCommentRestaurantDataStorage
}

type CommentRestaurantId int64

type CommentRestaurantUseCase struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type CommentsRestaurantUseCase struct {
	Comment []CommentRestaurantUseCase
}
