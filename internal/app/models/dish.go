package models

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
