package restaurantsHandler

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
// 	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRestaurantsHandler_GetAllRestaurants(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

// 	mockUCase := new(mock.RestaurantsUcase)

// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/restaurants", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := NewRestaurantsHandler(mockUCase, staticManager)

// 	json := `[{"id":1,"restName":"Name","imgPath":"http://localhost/static/restaurants/ImagePath","slug":"slug","price":1,"rating":4.5,"timeToDeliver":"2-3"}]`

// 	err = handler.GetAllRestaurantsMain(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestRestaurantsHandler_GetAllRestaurants_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

// 	mockUCase := new(mock.RestaurantsUcaseErr)
// 	mockLogger := new(mockLogger.Logger)

// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/restaurants", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := NewRestaurantsHandler(mockUCase, staticManager)

// 	err = handler.GetAllRestaurantsMain(c)
// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }

// func TestRestaurantsHandler_GetRestaurantsByCategory(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

// 	mockUCase := new(mock.RestaurantsUcase)

// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/restaurants?category=Суши", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := NewRestaurantsHandler(mockUCase, staticManager)

// 	json := `[{"id":1,"restName":"Name","imgPath":"http://localhost/static/restaurants/ImagePath","slug":"slug","price":1,"rating":4.5,"timeToDeliver":"2-3"}]`

// 	err = handler.GetAllRestaurantsMain(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestRestaurantsHandler_GetRestaurantsByCategory_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

// 	mockUCase := new(mock.RestaurantsUcaseErr)
// 	mockLogger := new(mockLogger.Logger)

// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/restaurants?category=Суши", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := NewRestaurantsHandler(mockUCase, staticManager)

// 	err = handler.GetAllRestaurantsMain(c)
// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }

// func TestRestaurantsHandler_GetRestaurantsByQuery(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

// 	mockUCase := new(mock.RestaurantsUcase)

// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/restaurants?q=Суши", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := NewRestaurantsHandler(mockUCase, staticManager)

// 	json := `[{"id":1,"restName":"Name","imgPath":"http://localhost/static/restaurants/ImagePath","slug":"slug","price":1,"rating":4.5,"timeToDeliver":"2-3"}]`

// 	err = handler.GetAllRestaurantsMain(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestRestaurantsHandler_GetRestaurantsByQuery_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

// 	mockUCase := new(mock.RestaurantsUcaseErr)
// 	mockLogger := new(mockLogger.Logger)

// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/restaurants?q=Суши", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := NewRestaurantsHandler(mockUCase, staticManager)

// 	err = handler.GetAllRestaurantsMain(c)
// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }
