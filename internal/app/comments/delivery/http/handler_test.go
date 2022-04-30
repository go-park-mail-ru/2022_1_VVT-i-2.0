package restaurantsHandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommentsHandler_GetRestaurantComments(t *testing.T) {
	id := "1"

	mockUCase := new(interfaces.CommentsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/comments/1"+id, strings.NewReader("1"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("id", id)
	handler := CommentsHandler{
		Usecase: mockUCase,
	}

	json := `[{"id":1,"restaurants":1,"userId":1,"commentText":"comment","commentRating":5}]`

	err = handler.GetRestaurantComments(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestRestaurantsHandler_AddCommentsRestaurantByRestaurants(t *testing.T) {

	commentRequest := models.AddCommentRestaurant{
		Restaurant:     data.CommentRestaurant.Restaurant,
		User_id:        data.CommentRestaurant.User_id,
		Comment_text:   data.CommentRestaurant.Comment_text,
		Comment_rating: data.CommentRestaurant.Comment_rating,
	}

	j, err := json.Marshal(commentRequest)

	mockUCase := new(interfaces.CommentsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/comment", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := CommentsHandler{
		Usecase: mockUCase,
	}

	json := `{"id":1,"restaurants":1,"userId":1,"commentText":"comment","commentRating":5}`

	err = handler.AddRestaurantComment(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}
