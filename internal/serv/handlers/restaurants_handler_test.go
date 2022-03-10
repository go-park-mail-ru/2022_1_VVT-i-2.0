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

type TestCase struct {
	Json       models.City
	Response   string
	StatusCode int
}

func TestRestaurants(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Json:       models.City{City: ""},
			Response:   `{"restaurants":[{"id":1,"imgPath":"http://178.154.229.61:8080/static/unsplash_HlNcigvUi4Q.png","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":2,"imgPath":"http://178.154.229.61:8080/static/pic.jpg","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":3,"imgPath":"http://178.154.229.61:8080/static/pic.jpg","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":4,"imgPath":"http://178.154.229.61:8080/static/pic.jpg","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8}],"auth":false,"city":"moscow"}`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			Json:       models.City{City: "voronezh"},
			Response:   `{"restaurants":[{"id":1,"imgPath":"http://178.154.229.61:8080/static/unsplash_HlNcigvUi4Q.png","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":2,"imgPath":"http://178.154.229.61:8080/static/pic.jpg","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":3,"imgPath":"http://178.154.229.61:8080/static/pic.jpg","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8},{"id":4,"imgPath":"http://178.154.229.61:8080/static/pic.jpg","restName":"Шоколадница","timeToDeliver":"20-35 мин","price":"550₽","rating":4.8}],"auth":false,"city":"voronezh"}`,
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
