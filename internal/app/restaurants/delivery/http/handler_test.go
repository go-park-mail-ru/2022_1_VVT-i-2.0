package restaurantsHandler

import (
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
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

	req, err := http.NewRequest(echo.GET, "restaurants", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	err = handler.GetAllRestaurants(c)
	require.NoError(t, err)
}

func TestRestaurantsHandler_GetDishesByRestaurants(t *testing.T) {
	slug := "my_slug"

	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.RestaurantsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "restaurant/" + slug, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("slug", slug)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	err = handler.GetDishesByRestaurants(c)
	require.NoError(t, err)
}

func TestNewRestaurantsHandler_GetCommentsRestaurantByRestaurants(t *testing.T) {
	id := "1"

	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.RestaurantsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "comments/" + id, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("id", id)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	err = handler.GetCommentsRestaurantByRestaurants(c)
	require.NoError(t, err)
}

func TestNewRestaurantsHandler_AddCommentsRestaurantByRestaurants(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.RestaurantsUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "comment/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	err = handler.GetCommentsRestaurantByRestaurants(c)
	require.NoError(t, err)
}
