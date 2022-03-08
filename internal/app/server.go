package serv

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *zap.SugaredLogger
}

func newServer(logger *zap.SugaredLogger) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

type Restaurant struct {
	ID       	 	int 	`json:"id"`
	ImagePath	 	string 	`json:"imgPath"`
	Name  		 	string 	`json:"restName"`
	TimeToDeliver 	string	`json:"timeToDeliver"`
	Price 			string 	`json:"price"`
	Rating 			float64 `json:"rating"`
}

type Answer struct {
	Restaurants []Restaurant `json:"restaurants"`
	Auth bool `json:"auth"`
	City string `json:"city"`
}

type City struct {
	City string `json:"city"`
}

var restaurant = []Restaurant{
	{ID: 1, ImagePath: "url", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 2, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 3, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 4, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	//{ID: 5, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	//{ID: 6, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
}

var answer = Answer {
	Restaurants: restaurant,
}



//json || по json я ставлю куку
//мы сначала смотри авторизазацию в контексте
//потом в куке
//москва

func restaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p City

	userId := r.Context().Value("keyUserId")
	var auth = false
	if userId != nil {
		fmt.Println("\nhello, %s", userId)
		auth = true
	} else {
		fmt.Println("\nhello, incognito")
	}
	answer.Auth = auth

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {

	} else {
		fmt.Fprintf(w, "Person: %+v\n", p)
	}


	city, err := r.Cookie("city")
	existCity := err != http.ErrNoCookie

	if existCity {
		fmt.Fprintln(w, "Welcome, "+city.Value)
		answer.City = city.Value
	} else {
		fmt.Fprintln(w, "You need to login")
	}

	vars := mux.Vars(r)
	if cookieCity, found  := vars["city"]; found {
		fmt.Fprintf(w, "City: %s\n\n", cookieCity)
		cookie := http.Cookie{
			Name:    "city",
			Value:   cookieCity,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, "City: %s\n\n", cookieCity)
	}



	result, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	w.Write(result)
}

func (s *server) configureRouter() {
	// TODO: set prefix "api/v1" anywhere

	noAuthRequiredRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRequiredRouter.HandleFunc("/restaurants", restaurants)
	// noAuthRequiredRouter.Use(s.AuthMiddleware)

	authRequiredRouter := s.router.PathPrefix("/auth").Subrouter()
	authRequiredRouter.HandleFunc("/h", hello)
	// authRequiredRouter.Use(s.RequiredAuthMiddleware)

	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println("hello")
}

// s.router.HandleFunc("/restorants/{city}/{page_num}", getRestaurants).Methods(http.MethodGet)
// getRestarants нужно вернуть список ресторанов в данном городе, установить куку города
// инфа обработчика: колво ресторанов на странице
