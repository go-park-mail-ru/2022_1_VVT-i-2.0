package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
)

type testCase struct {
	Json       models.City
	Response   string
	StatusCode int
}

func TestRestaurants(t *testing.T) {
	cases := []testCase{
		testCase{
			Json:       models.City{City: ""},
			Response:   `{"restaurants":[{"id":1,"imgPath":"http://tavide.xyz:8080/static/unsplash_HlNcigvUi4Q.png","restName":"Шоколадница","timeToDeliver":"20-45 мин","price":"650₽","rating":4.8},{"id":2,"imgPath":"http://tavide.xyz:8080/static/smekalca_food.png","restName":"Smekalca FooD","timeToDeliver":"20-35 мин","price":"570₽","rating":4.7},{"id":3,"imgPath":"http://tavide.xyz:8080/static/subway.png","restName":"Subway","timeToDeliver":"20-55 мин","price":"1050₽","rating":4.6},{"id":4,"imgPath":"http://tavide.xyz:8080/static/shaurma.png","restName":"Шаурма","timeToDeliver":"25-35 мин","price":"350₽","rating":4.5},{"id":5,"imgPath":"http://tavide.xyz:8080/static/mac.png","restName":"Макдональдс","timeToDeliver":"10-35 мин","price":"650₽","rating":4.4},{"id":6,"imgPath":"http://tavide.xyz:8080/static/KFC.png","restName":"KFC","timeToDeliver":"20-35 мин","price":"550₽","rating":4.3},{"id":7,"imgPath":"http://tavide.xyz:8080/static/BK.png","restName":"Burger King","timeToDeliver":"20-35 мин","price":"770₽","rating":4.2},{"id":8,"imgPath":"http://tavide.xyz:8080/static/terem.png","restName":"Теремок","timeToDeliver":"25-35 мин","price":"665₽","rating":4.1},{"id":9,"imgPath":"http://tavide.xyz:8080/static/zotman.png","restName":"Zotmann Pizza","timeToDeliver":"20-55 мин","price":"2340₽","rating":4},{"id":10,"imgPath":"http://tavide.xyz:8080/static/tuktuk.png","restName":"Tuk Tuk","timeToDeliver":"20-35 мин","price":"1000₽","rating":4.8},{"id":11,"imgPath":"http://tavide.xyz:8080/static/Bo.png","restName":"BO","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":12,"imgPath":"http://tavide.xyz:8080/static/paple.png","restName":"Крошка картошка","timeToDeliver":"20-65 мин","price":"750₽","rating":5},{"id":13,"imgPath":"http://tavide.xyz:8080/static/yaki.png","restName":"Якитория","timeToDeliver":"30-35 мин","price":"850₽","rating":4.8},{"id":14,"imgPath":"http://tavide.xyz:8080/static/dad.png","restName":"Мама джанс","timeToDeliver":"35-45 мин","price":"950₽","rating":4.7},{"id":15,"imgPath":"http://tavide.xyz:8080/static/carlamov.png","restName":"Варламов.сесть","timeToDeliver":"25-35 мин","price":"550₽","rating":4.8},{"id":16,"imgPath":"http://tavide.xyz:8080/static/allo.png","restName":"Алло!Пицца","timeToDeliver":"20-50 мин","price":"450₽","rating":4.6},{"id":17,"imgPath":"http://tavide.xyz:8080/static/fo89.png","restName":"Fo 98","timeToDeliver":"20-50 мин","price":"560₽","rating":4.7},{"id":18,"imgPath":"http://tavide.xyz:8080/static/pizzaexp.png","restName":"Pizza Express 25/8","timeToDeliver":"20-35 мин","price":"656₽","rating":4.8},{"id":19,"imgPath":"http://tavide.xyz:8080/static/tanuki.png","restName":"Tanuki","timeToDeliver":"20-40 мин","price":"770₽","rating":4.7},{"id":20,"imgPath":"http://tavide.xyz:8080/static/chay.png","restName":"Чайона №2","timeToDeliver":"20-35 мин","price":"777₽","rating":4.6},{"id":21,"imgPath":"http://tavide.xyz:8080/static/sakura.png","restName":"Sakura","timeToDeliver":"20-55 мин","price":"770₽","rating":4.8}],"auth":false,"city":"moscow"}`,
			StatusCode: http.StatusOK,
		},
		testCase{
			Json:       models.City{City: "moscow"},
			Response:   `{"restaurants":[{"id":1,"imgPath":"http://tavide.xyz:8080/static/unsplash_HlNcigvUi4Q.png","restName":"Шоколадница","timeToDeliver":"20-45 мин","price":"650₽","rating":4.8},{"id":2,"imgPath":"http://tavide.xyz:8080/static/smekalca_food.png","restName":"Smekalca FooD","timeToDeliver":"20-35 мин","price":"570₽","rating":4.7},{"id":3,"imgPath":"http://tavide.xyz:8080/static/subway.png","restName":"Subway","timeToDeliver":"20-55 мин","price":"1050₽","rating":4.6},{"id":4,"imgPath":"http://tavide.xyz:8080/static/shaurma.png","restName":"Шаурма","timeToDeliver":"25-35 мин","price":"350₽","rating":4.5},{"id":5,"imgPath":"http://tavide.xyz:8080/static/mac.png","restName":"Макдональдс","timeToDeliver":"10-35 мин","price":"650₽","rating":4.4},{"id":6,"imgPath":"http://tavide.xyz:8080/static/KFC.png","restName":"KFC","timeToDeliver":"20-35 мин","price":"550₽","rating":4.3},{"id":7,"imgPath":"http://tavide.xyz:8080/static/BK.png","restName":"Burger King","timeToDeliver":"20-35 мин","price":"770₽","rating":4.2},{"id":8,"imgPath":"http://tavide.xyz:8080/static/terem.png","restName":"Теремок","timeToDeliver":"25-35 мин","price":"665₽","rating":4.1},{"id":9,"imgPath":"http://tavide.xyz:8080/static/zotman.png","restName":"Zotmann Pizza","timeToDeliver":"20-55 мин","price":"2340₽","rating":4},{"id":10,"imgPath":"http://tavide.xyz:8080/static/tuktuk.png","restName":"Tuk Tuk","timeToDeliver":"20-35 мин","price":"1000₽","rating":4.8},{"id":11,"imgPath":"http://tavide.xyz:8080/static/Bo.png","restName":"BO","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":12,"imgPath":"http://tavide.xyz:8080/static/paple.png","restName":"Крошка картошка","timeToDeliver":"20-65 мин","price":"750₽","rating":5},{"id":13,"imgPath":"http://tavide.xyz:8080/static/yaki.png","restName":"Якитория","timeToDeliver":"30-35 мин","price":"850₽","rating":4.8},{"id":14,"imgPath":"http://tavide.xyz:8080/static/dad.png","restName":"Мама джанс","timeToDeliver":"35-45 мин","price":"950₽","rating":4.7},{"id":15,"imgPath":"http://tavide.xyz:8080/static/carlamov.png","restName":"Варламов.сесть","timeToDeliver":"25-35 мин","price":"550₽","rating":4.8},{"id":16,"imgPath":"http://tavide.xyz:8080/static/allo.png","restName":"Алло!Пицца","timeToDeliver":"20-50 мин","price":"450₽","rating":4.6},{"id":17,"imgPath":"http://tavide.xyz:8080/static/fo89.png","restName":"Fo 98","timeToDeliver":"20-50 мин","price":"560₽","rating":4.7},{"id":18,"imgPath":"http://tavide.xyz:8080/static/pizzaexp.png","restName":"Pizza Express 25/8","timeToDeliver":"20-35 мин","price":"656₽","rating":4.8},{"id":19,"imgPath":"http://tavide.xyz:8080/static/tanuki.png","restName":"Tanuki","timeToDeliver":"20-40 мин","price":"770₽","rating":4.7},{"id":20,"imgPath":"http://tavide.xyz:8080/static/chay.png","restName":"Чайона №2","timeToDeliver":"20-35 мин","price":"777₽","rating":4.6},{"id":21,"imgPath":"http://tavide.xyz:8080/static/sakura.png","restName":"Sakura","timeToDeliver":"20-55 мин","price":"770₽","rating":4.8}],"auth":false,"city":"moscow"}`,
			StatusCode: http.StatusOK,
		},
		testCase{
			Json:       models.City{City: "voronezh"},
			Response:   `{"restaurants":[{"id":1,"imgPath":"http://tavide.xyz:8080/static/unsplash_HlNcigvUi4Q.png","restName":"Шоколадница","timeToDeliver":"20-45 мин","price":"650₽","rating":4.8},{"id":2,"imgPath":"http://tavide.xyz:8080/static/smekalca_food.png","restName":"Smekalca FooD","timeToDeliver":"20-35 мин","price":"570₽","rating":4.7},{"id":3,"imgPath":"http://tavide.xyz:8080/static/subway.png","restName":"Subway","timeToDeliver":"20-55 мин","price":"1050₽","rating":4.6},{"id":4,"imgPath":"http://tavide.xyz:8080/static/shaurma.png","restName":"Шаурма","timeToDeliver":"25-35 мин","price":"350₽","rating":4.5},{"id":5,"imgPath":"http://tavide.xyz:8080/static/mac.png","restName":"Макдональдс","timeToDeliver":"10-35 мин","price":"650₽","rating":4.4},{"id":6,"imgPath":"http://tavide.xyz:8080/static/KFC.png","restName":"KFC","timeToDeliver":"20-35 мин","price":"550₽","rating":4.3},{"id":7,"imgPath":"http://tavide.xyz:8080/static/BK.png","restName":"Burger King","timeToDeliver":"20-35 мин","price":"770₽","rating":4.2},{"id":8,"imgPath":"http://tavide.xyz:8080/static/terem.png","restName":"Теремок","timeToDeliver":"25-35 мин","price":"665₽","rating":4.1},{"id":9,"imgPath":"http://tavide.xyz:8080/static/zotman.png","restName":"Zotmann Pizza","timeToDeliver":"20-55 мин","price":"2340₽","rating":4},{"id":10,"imgPath":"http://tavide.xyz:8080/static/tuktuk.png","restName":"Tuk Tuk","timeToDeliver":"20-35 мин","price":"1000₽","rating":4.8},{"id":11,"imgPath":"http://tavide.xyz:8080/static/Bo.png","restName":"BO","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":12,"imgPath":"http://tavide.xyz:8080/static/paple.png","restName":"Крошка картошка","timeToDeliver":"20-65 мин","price":"750₽","rating":5},{"id":13,"imgPath":"http://tavide.xyz:8080/static/yaki.png","restName":"Якитория","timeToDeliver":"30-35 мин","price":"850₽","rating":4.8},{"id":14,"imgPath":"http://tavide.xyz:8080/static/dad.png","restName":"Мама джанс","timeToDeliver":"35-45 мин","price":"950₽","rating":4.7},{"id":15,"imgPath":"http://tavide.xyz:8080/static/carlamov.png","restName":"Варламов.сесть","timeToDeliver":"25-35 мин","price":"550₽","rating":4.8},{"id":16,"imgPath":"http://tavide.xyz:8080/static/allo.png","restName":"Алло!Пицца","timeToDeliver":"20-50 мин","price":"450₽","rating":4.6},{"id":17,"imgPath":"http://tavide.xyz:8080/static/fo89.png","restName":"Fo 98","timeToDeliver":"20-50 мин","price":"560₽","rating":4.7},{"id":18,"imgPath":"http://tavide.xyz:8080/static/pizzaexp.png","restName":"Pizza Express 25/8","timeToDeliver":"20-35 мин","price":"656₽","rating":4.8},{"id":19,"imgPath":"http://tavide.xyz:8080/static/tanuki.png","restName":"Tanuki","timeToDeliver":"20-40 мин","price":"770₽","rating":4.7},{"id":20,"imgPath":"http://tavide.xyz:8080/static/chay.png","restName":"Чайона №2","timeToDeliver":"20-35 мин","price":"777₽","rating":4.6},{"id":21,"imgPath":"http://tavide.xyz:8080/static/sakura.png","restName":"Sakura","timeToDeliver":"20-55 мин","price":"770₽","rating":4.8}],"auth":false,"city":"voronezh"}`,
			StatusCode: http.StatusOK,
		},
		testCase{
			Json:       models.City{City: "new york"},
			Response:   `{"restaurants":[{"id":1,"imgPath":"http://tavide.xyz:8080/static/unsplash_HlNcigvUi4Q.png","restName":"Шоколадница","timeToDeliver":"20-45 мин","price":"650₽","rating":4.8},{"id":2,"imgPath":"http://tavide.xyz:8080/static/smekalca_food.png","restName":"Smekalca FooD","timeToDeliver":"20-35 мин","price":"570₽","rating":4.7},{"id":3,"imgPath":"http://tavide.xyz:8080/static/subway.png","restName":"Subway","timeToDeliver":"20-55 мин","price":"1050₽","rating":4.6},{"id":4,"imgPath":"http://tavide.xyz:8080/static/shaurma.png","restName":"Шаурма","timeToDeliver":"25-35 мин","price":"350₽","rating":4.5},{"id":5,"imgPath":"http://tavide.xyz:8080/static/mac.png","restName":"Макдональдс","timeToDeliver":"10-35 мин","price":"650₽","rating":4.4},{"id":6,"imgPath":"http://tavide.xyz:8080/static/KFC.png","restName":"KFC","timeToDeliver":"20-35 мин","price":"550₽","rating":4.3},{"id":7,"imgPath":"http://tavide.xyz:8080/static/BK.png","restName":"Burger King","timeToDeliver":"20-35 мин","price":"770₽","rating":4.2},{"id":8,"imgPath":"http://tavide.xyz:8080/static/terem.png","restName":"Теремок","timeToDeliver":"25-35 мин","price":"665₽","rating":4.1},{"id":9,"imgPath":"http://tavide.xyz:8080/static/zotman.png","restName":"Zotmann Pizza","timeToDeliver":"20-55 мин","price":"2340₽","rating":4},{"id":10,"imgPath":"http://tavide.xyz:8080/static/tuktuk.png","restName":"Tuk Tuk","timeToDeliver":"20-35 мин","price":"1000₽","rating":4.8},{"id":11,"imgPath":"http://tavide.xyz:8080/static/Bo.png","restName":"BO","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":12,"imgPath":"http://tavide.xyz:8080/static/paple.png","restName":"Крошка картошка","timeToDeliver":"20-65 мин","price":"750₽","rating":5},{"id":13,"imgPath":"http://tavide.xyz:8080/static/yaki.png","restName":"Якитория","timeToDeliver":"30-35 мин","price":"850₽","rating":4.8},{"id":14,"imgPath":"http://tavide.xyz:8080/static/dad.png","restName":"Мама джанс","timeToDeliver":"35-45 мин","price":"950₽","rating":4.7},{"id":15,"imgPath":"http://tavide.xyz:8080/static/carlamov.png","restName":"Варламов.сесть","timeToDeliver":"25-35 мин","price":"550₽","rating":4.8},{"id":16,"imgPath":"http://tavide.xyz:8080/static/allo.png","restName":"Алло!Пицца","timeToDeliver":"20-50 мин","price":"450₽","rating":4.6},{"id":17,"imgPath":"http://tavide.xyz:8080/static/fo89.png","restName":"Fo 98","timeToDeliver":"20-50 мин","price":"560₽","rating":4.7},{"id":18,"imgPath":"http://tavide.xyz:8080/static/pizzaexp.png","restName":"Pizza Express 25/8","timeToDeliver":"20-35 мин","price":"656₽","rating":4.8},{"id":19,"imgPath":"http://tavide.xyz:8080/static/tanuki.png","restName":"Tanuki","timeToDeliver":"20-40 мин","price":"770₽","rating":4.7},{"id":20,"imgPath":"http://tavide.xyz:8080/static/chay.png","restName":"Чайона №2","timeToDeliver":"20-35 мин","price":"777₽","rating":4.6},{"id":21,"imgPath":"http://tavide.xyz:8080/static/sakura.png","restName":"Sakura","timeToDeliver":"20-55 мин","price":"770₽","rating":4.8}],"auth":false,"city":"moscow"}`,
			StatusCode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://example.com/api/v1/restaurants"
		payloadBuf := new(bytes.Buffer)
		if item.Json.City != "" {
			json.NewEncoder(payloadBuf).Encode(item.Json)
			fmt.Printf("%s\n", item.Json)
			fmt.Printf("%s\n", payloadBuf)
		} else {
			json.NewEncoder(payloadBuf).Encode("")
		}
		req := httptest.NewRequest("GET", url, payloadBuf)
		w := httptest.NewRecorder()

		RestaurantsHandler(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}
