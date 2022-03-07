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

var u = []Restaurant{
	{ID: 1, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 2, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 3, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 4, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 5, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 6, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city, err := r.Cookie("session_id")
	existCity := err != http.ErrNoCookie

	if existCity {
		fmt.Fprintln(w, "Welcome, "+city.Value)
	} else {
		fmt.Fprintln(w, "You need to login")
	}

	auth := true
	if auth {
		fmt.Fprintf(w, "auth: %b\n", auth)
	} else {
		fmt.Fprintf(w, "auth: %b\n", auth)
	}


	vars := mux.Vars(r)
	if cookieCity, found  := vars["city"]; found {
		cookie := http.Cookie{
			Name:    "city",
			Value:   cookieCity,
			Secure:   true,
			HttpOnly: true,

		}
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, "City: %s\n\n", city)
	}



	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	w.Write(result)
}

func (s *server) configureRouter() {
	// TODO: set prefix "api/v1" anywhere

	noAuthRequiredRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRequiredRouter.HandleFunc("/restaurants", restaurants)
	noAuthRequiredRouter.HandleFunc("/restaurants/{city}", restaurants)
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
