package orderHandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestOrderHandler_GetOrder(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUCase, staticManager)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("orderId")
	c.SetParamValues("1")
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.GetUserOrder(c)

	json := `{"orderNumber":1,"address":"","date":"01.01.2021","totalPrice":100,"restName":"RestName","restSlug":"","status":"Получен","cart":[{"name":"name","description":"description","count":10,"price":10,"calories":100,"weight":50,"imgPath":"http://localhost/static/dishes/img.png"}]}`

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestOrderHandler_GetOrder_Err(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewOrderHandler(mockUCase, staticManager)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})

	c.SetParamNames("orderId")
	c.SetParamValues("1")
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.GetUserOrder(c)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestOrderHandler_CreateOrder(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUCase, staticManager)
	e := echo.New()

	reqBody := models.OrderReq{
		Address: "Москва, Измайловский проспект, 73/2",
		Comment: "comment",
		Cart:    []models.OrderPosition{{Id: 1, Count: 2}},
	}

	jsonReq, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(echo.POST, "/order", strings.NewReader(string(jsonReq)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	json := `{"OrderId":1}
`
	err := handler.CreateOrder(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestOrderHandler_CreateOrder_Err(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewOrderHandler(mockUCase, staticManager)
	e := echo.New()

	reqBody := models.OrderReq{
		Address: "Москва, Измайловский проспект, 73/2",
		Comment: "comment",
		Cart:    []models.OrderPosition{{Id: 1, Count: 2}},
	}

	jsonReq, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(echo.POST, "/order", strings.NewReader(string(jsonReq)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err := handler.CreateOrder(c)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestOrderHandler_GetOrders(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUCase, staticManager)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.GetUserOrders(c)

	json := `{"orderList":[{"orderNumber":1,"date":"01.01.2021","totalPrice":100,"restName":"RestName","status":"Получен"}]}`

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestOrderHandler_GetOrders_Err(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewOrderHandler(mockUCase, staticManager)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.GetUserOrders(c)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestOrderHandler_GetOrderStatuses(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUCase, staticManager)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order_statuses", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.GetUserOrderStatuses(c)

	json := `{"statuses":[{"id":1,"status":"Получен"}]}`

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestOrderHandler_GetOrderStatuses_Err(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")
	mockUCase := new(mock.OrderUcaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewOrderHandler(mockUCase, staticManager)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/order_statuses", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.GetUserOrderStatuses(c)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}
