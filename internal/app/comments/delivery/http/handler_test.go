package restaurantsHandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCommentsHandler_GetRestaurantComments(t *testing.T) {
	id := "1"

	mockUCase := new(mock.CommentsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/comments/1"+id, strings.NewReader("1"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("id", id)
	handler := CommentsHandler{
		Usecase: mockUCase,
	}

	json := `[{"author":"author","text":"text","stars":4,"date":"date"}]`

	err = handler.GetRestaurantComments(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestRestaurantsHandler_AddCommentsRestaurantByRestaurants(t *testing.T) {

	commentRequest := models.AddCommentRestaurant{
		Slug:          "slug",
		CommentText:   "text",
		CommentRating: 5,
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
	handler := CommentsHandler{
		Usecase: mockUCase,
	}

	json := `{"restaurants_id":1,"author":"author","text":"text","stars":4,"date":"date"}`

	err = handler.AddRestaurantComment(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestCommentsHandler_GetRestaurantComments_Err(t *testing.T) {
	id := "1"

	mockUCase := new(mock.CommentsUsecaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewCommentsHandler(mockUCase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/comments/1"+id, strings.NewReader("1"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	c.Set("id", id)

	err = handler.GetRestaurantComments(c)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestRestaurantsHandler_AddCommentsRestaurant_Err(t *testing.T) {

	commentRequest := models.AddCommentRestaurant{
		Slug:          "slug",
		CommentText:   "text",
		CommentRating: 5,
	}

	j, _ := json.Marshal(commentRequest)

	mockUCase := new(mock.CommentsUsecaseErr)
	mockLogger := new(mockLogger.Logger)
	handler := NewCommentsHandler(mockUCase)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/comment", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	err = handler.AddRestaurantComment(c)

	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}
