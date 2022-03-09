package serv

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Restaurant struct {
	ID            int     `json:"id"`
	ImagePath     string  `json:"imgPath"`
	Name          string  `json:"restName"`
	TimeToDeliver string  `json:"timeToDeliver"`
	Price         string  `json:"price"`
	Rating        float64 `json:"rating"`
}

type Answer struct {
	Restaurants []Restaurant `json:"restaurants"`
	Auth        bool         `json:"auth"`
	City        string       `json:"city"`
}

type City struct {
	City string `json:"city"`
}

var restaurant = []Restaurant{
	{ID: 1, ImagePath: "unsplash_HlNcigvUi4Q.png", Name: "Шоколадница", TimeToDeliver: "20-45 мин", Price: "650₽", Rating: 4.8},
	{ID: 2, ImagePath: "smekalca_food.png", Name: "Smekalca FooD", TimeToDeliver: "20-35 мин", Price: "570₽", Rating: 4.7},
	{ID: 3, ImagePath: "subway.png", Name: "Subway", TimeToDeliver: "20-55 мин", Price: "1050₽", Rating: 4.6},
	{ID: 4, ImagePath: "shaurma.png", Name: "Шаурма", TimeToDeliver: "25-35 мин", Price: "350₽", Rating: 4.5},
	{ID: 5, ImagePath: "mac.png", Name: "Макдональдс", TimeToDeliver: "10-35 мин", Price: "650₽", Rating: 4.4},
	{ID: 6, ImagePath: "KFC.png", Name: "KFC", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.3},
	{ID: 7, ImagePath: "BK.png", Name: "Burger King", TimeToDeliver: "20-35 мин", Price: "770₽", Rating: 4.2},
	{ID: 8, ImagePath: "terem.png", Name: "Теремок", TimeToDeliver: "25-35 мин", Price: "665₽", Rating: 4.1},
	{ID: 9, ImagePath: "zotman.png", Name: "Zotmann Pizza", TimeToDeliver: "20-55 мин", Price: "2340₽", Rating: 4.0},
	{ID: 10, ImagePath: "tuktuk.png", Name: "Tuk Tuk", TimeToDeliver: "20-35 мин", Price: "1000₽", Rating: 4.8},
	{ID: 11, ImagePath: "Bo.png", Name: "BO", TimeToDeliver: "20-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 12, ImagePath: "paple.png", Name: "Крошка картошка", TimeToDeliver: "20-65 мин", Price: "750₽", Rating: 5.0},
	{ID: 13, ImagePath: "yaki.png", Name: "Якитория", TimeToDeliver: "30-35 мин", Price: "850₽", Rating: 4.8},
	{ID: 14, ImagePath: "dad.png", Name: "Мама джанс", TimeToDeliver: "35-45 мин", Price: "950₽", Rating: 4.7},
	{ID: 15, ImagePath: "carlamov.png", Name: "Варламов.сесть", TimeToDeliver: "25-35 мин", Price: "550₽", Rating: 4.8},
	{ID: 16, ImagePath: "allo.png", Name: "Алло!Пицца", TimeToDeliver: "20-50 мин", Price: "450₽", Rating: 4.6},
	{ID: 17, ImagePath: "fo89.png", Name: "Fo 98", TimeToDeliver: "20-50 мин", Price: "560₽", Rating: 4.7},
	{ID: 18, ImagePath: "pizzaexp.png", Name: "Pizza Express 25/8", TimeToDeliver: "20-35 мин", Price: "656₽", Rating: 4.8},
	{ID: 19, ImagePath: "tanuki.png", Name: "Tanuki", TimeToDeliver: "20-40 мин", Price: "770₽", Rating: 4.7},
	{ID: 20, ImagePath: "chay.png", Name: "Чайона №2", TimeToDeliver: "20-35 мин", Price: "777₽", Rating: 4.6},
	{ID: 21, ImagePath: "sakura.png", Name: "Sakura", TimeToDeliver: "20-55 мин", Price: "770₽", Rating: 4.8},

}

//json || по json я ставлю куку
//мы сначала смотри авторизазацию в контексте
//потом в куке
//москва

func getCityFromDb(userId uint64) string {
	return string("moscow")
}

func workWithURL(rest []Restaurant) []Restaurant {
	var restaurant []Restaurant
	domen := "tavide.xyz"
	port := "8080"
	directory := "static"
	buff := ""
	buffStruct := Restaurant{}
	for i, _ := range rest {
		buffStruct = rest[i]
		buff = rest[i].ImagePath
		mySuperString := "http://" + domen + ":" + port + "/" + directory + "/" + buff
		buffStruct.ImagePath = mySuperString
		restaurant = append(restaurant, buffStruct)
	}
	return restaurant
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jsonCity City
	var answer = Answer{}

	var userId uint64
	if r.Context().Value(keyUserId) != nil {
		userId = uint64(r.Context().Value(keyUserId).(ctxUserId))
	}

	var auth = userId != 0
	answer.Auth = auth

	err := json.NewDecoder(r.Body).Decode(&jsonCity)
	if err != nil {
		if auth {
			answer.City = getCityFromDb(uint64(userId))
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
			Name:     "city",
			Value:    jsonCity.City,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
	}

	answer.Restaurants = workWithURL(restaurant)

	result, err := json.Marshal(answer)
	if err != nil {
		fmt.Printf("Marshal error\n")
	}
	w.Write(result)
}
