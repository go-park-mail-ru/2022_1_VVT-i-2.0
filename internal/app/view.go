package serv

import (
	"encoding/json"
	"fmt"
	"net/http"
)


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
	{ID: 2, ImagePath: "ulr", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 3, ImagePath: "url", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 4, ImagePath: "url", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	//{ID: 5, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	//{ID: 6, ImagePath: "", Name: "Шоколадница", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
}

//json || по json я ставлю куку
//мы сначала смотри авторизазацию в контексте
//потом в куке
//москва

func workWithURL(rest []Restaurant) {
	domen := "127.0.0.1"
	port := "8080"
	directory := "static"
	buff := ""
	for i, _ := range rest {
		buff = rest[i].ImagePath
		mySuperString := "http://" + domen + ":" + port + "/" + directory + "/" + buff
		rest[i].ImagePath = mySuperString
	}
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jsonCity City
	var answer = Answer{}

	user := r.Context().Value(keyUser).(ctxStruct).user
	var auth = false
	if user != nil {
		fmt.Println("\nhello, %s", user)
		auth = true
	} else {
		fmt.Println("\nhello, incognito")
	}
	answer.Auth = auth

	err := json.NewDecoder(r.Body).Decode(&jsonCity)
	if err != nil {
		if auth {
			answer.City = user.address
			fmt.Printf("город выставлен по контексту\n")
		} else {
			city, err := r.Cookie("city")
			existCity := err != http.ErrNoCookie

			if existCity {
				fmt.Printf("%s\n", "город выставлен по cookie")
				fmt.Printf("Welcome %s\n", city.Value)
				answer.City = city.Value
			} else {
				fmt.Printf("%s\n", "город выставлен по умолчанию")
				fmt.Printf("%s\n", "You need to login")
				answer.City = "moscow"
			}
		}
	} else {
		fmt.Printf("%s\n", "город выставлен по json")
		cookie := http.Cookie{
			Name:    "city",
			Value:   jsonCity.City,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
	}

	workWithURL(restaurant)
	answer.Restaurants = restaurant


	result, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	w.Write(result)
}
