package models

type UserId uint64

type Restaurant struct {
	ID            int     `json:"id"`
	ImagePath     string  `json:"imgPath"`
	Name          string  `json:"restName"`
	TimeToDeliver string  `json:"timeToDeliver"`
	Price         string  `json:"price"`
	Rating        float64 `json:"rating"`
}

type RestaurantsResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
	Auth        bool         `json:"auth"`
	City        string       `json:"city"`
}

type City struct {
	City string `json:"city"`
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username,omitempty"`
	UserAddr string `json:"userAddress,omitempty"`
	Err      string `json:"error,omitempty"`
}

type UserDataStruct struct {
	Id      UserId
	Name    string
	Address string
}

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Username string `json:"name"`
}

type RegisterResponse struct {
	Username string `json:"username,omitempty"`
	UserAddr string `json:"userAddress,omitempty"`
	Err      string `json:"error,omitempty"`
}
