package restaurantsHandler

import (
	"encoding/json"
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRestaurantsHandler_GetAllRestaurants(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)


	mockUCase := new(interfaces.RestaurantsUsecase)

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/restaurants", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	json := `[{"id":1,"restName":"Name","city":"City","address":"Address","imgPath":"http://127.0.0.1:8080/static/static/imgPath","slug":"slug","min_price":1,"price":1,"rating":1,"timeToDeliver":"25 - 30"}]`

	err = handler.GetAllRestaurants(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestRestaurantsHandler_GetDishesByRestaurants(t *testing.T) {
	slug := "my_slug"

	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.RestaurantsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/restaurant/my_slug", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("slug", slug)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	json := `{"id":1,"restName":"Name","city":"City","address":"Address","imgPath":"restaurants/imgPath","slug":"slug","minPrice":1,"avgPrice":1,"rating":1,"timeToDeliver":"25-30","dishes":[{"id":1,"restaurany":1,"productName":"Name","description":"Description","imgPath":"dishes/imgPath","info":1,"price":1,"weight":1}]}`

	err = handler.GetDishesByRestaurants(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestRestaurantsHandler_GetCommentsRestaurantByRestaurants(t *testing.T) {
	id := "1"

	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.RestaurantsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/comments/1" + id, strings.NewReader("1"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("id", id)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	json := `[{"id":1,"restaurants":1,"userId":1,"commentText":"comment","commentRating":5}]`

	err = handler.GetCommentsRestaurantByRestaurants(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestRestaurantsHandler_AddCommentsRestaurantByRestaurants(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	commentRequest := models.AddCommentRestaurant{
		Restaurant: 	data.CommentRestaurant.Restaurant,
		User_id: 		data.CommentRestaurant.User_id,
		Comment_text: 	data.CommentRestaurant.Comment_text,
		Comment_rating:	data.CommentRestaurant.Comment_rating,
	}

	j, err := json.Marshal(commentRequest)

	mockUCase := new(interfaces.RestaurantsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/comment", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	json := `{"id":1,"restaurants":1,"userId":1,"commentText":"comment","commentRating":5}`

	err = handler.AddCommentsRestaurantByRestaurants(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}
