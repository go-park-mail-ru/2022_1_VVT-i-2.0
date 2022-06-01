package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCommentsHandler_GetRestaurantComments(t *testing.T) {
	slug := "1"

	mockUCase := new(mock.CommentsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/comments/1", strings.NewReader("1"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("slug")
	c.SetParamValues(slug)
	handler := RestaurantReviewsHandler{
		Ucase: mockUCase,
	}

	json := `[{"author":"author","text":"text","stars":4,"date":"date"}]`

	err = handler.GetRestaurantComments(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestRestaurantsHandler_AddCommentsRestaurantByRestaurants(t *testing.T) {

	commentRequest := models.AddRestaurantReviewReq{
		Slug:   "slug",
		Text:   "text",
		Rating: 5,
	}

	j, _ := json.Marshal(commentRequest)

	mockUCase := new(mock.CommentsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/comment", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})
	handler := RestaurantReviewsHandler{
		Ucase: mockUCase,
	}

	json := `{"author":"author","text":"text","stars":4,"date":"date"}`

	err = handler.AddRestaurantReview(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestCommentsHandler_GetRestaurantComments_WithOutSlug(t *testing.T) {
	mockUCase := new(mock.CommentsUsecaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewRestaurantReviewsHandler(mockUCase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/comments/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})

	err = handler.GetRestaurantComments(c)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, c.Response().Status)
}

func TestRestaurantsHandler_AddCommentsRestaurant_Err(t *testing.T) {

	commentRequest := models.AddRestaurantReviewReq{
		Slug:   "slug",
		Text:   "text",
		Rating: 5,
	}

	j, _ := json.Marshal(commentRequest)

	mockUCase := new(mock.CommentsUsecaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewRestaurantReviewsHandler(mockUCase)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/comment", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.AddRestaurantReview(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}
