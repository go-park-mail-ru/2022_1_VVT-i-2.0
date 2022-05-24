package restaurantsHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDishesHandler_GetRestaurantDishes(t *testing.T) {
	slug := "my_slug"
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

	mockUCase := new(mock.DishesUcase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/restaurant", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("slug")
	c.SetParamValues(slug)
	handler := NewDishesHandler(mockUCase, staticManager)

	json := `{"id":1,"restName":"Name","imgPath":"http://localhost/static/restaurants/ImagePath","slug":"Slug","minPrice":1,"rating":4.5,"timeToDeliver":"2-3","reviewCount":2,"dishes":[{"id":1,"category":0,"restaurant":1,"productName":"Name","description":"Description","imgPath":"http://localhost/static/dishes/DishImagePath","info":200,"price":10,"weight":100}],"categories":[{"category":"1","dishes":[1]}]}`

	err = handler.GetDishesByRestaurants(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestDishesHandler_GetRestaurantDishes_Err(t *testing.T) {
	slug := "my_slug"
	staticManager := localStaticManager.NewLocalFileManager("http:/static", "/static")

	mockUCase := new(mock.DishesUcaseErr)
	mockLogger := new(mockLogger.Logger)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/restaurant/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("slug")
	c.SetParamValues(slug)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})

	handler := NewDishesHandler(mockUCase, staticManager)

	err = handler.GetDishesByRestaurants(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestDishesHandler_GetRestaurantDishes_EmptySlug(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http:/static", "/static")

	mockUCase := new(mock.DishesUcaseErr)
	mockLogger := new(mockLogger.Logger)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/restaurant/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})

	handler := NewDishesHandler(mockUCase, staticManager)

	err = handler.GetDishesByRestaurants(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, c.Response().Status)
}
