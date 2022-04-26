package restaurantsHandler

import (
	"github.com/bxcodec/faker"
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	interfaces "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	mockRestaurant := &models.RestaurantUsecase{}
	err = faker.FakeData(&mockRestaurant)
	assert.NoError(t, err)
	mockRestaurants := &models.RestaurantsUsecase{}
	mockRestaurants.Restaurants = append(mockRestaurants.Restaurants, *mockRestaurant)

	mockUCase := new(interfaces.ArticleUsecase)
	mockUCase.On("GetAllRestaurants", mock.Anything).Return(mockRestaurants)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/restaurants", strings.NewReader(""))
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

	mockRestaurant := &models.RestaurantUsecase{}
	err = faker.FakeData(&mockRestaurant)
	assert.NoError(t, err)

	mockUCase := new(interfaces.ArticleUsecase)
	mockUCase.On("GetRestaurantBySluf", mock.Anything, "slug").Return(mockRestaurant)

	mockDish := &models.DishUseCase{}
	err = faker.FakeData(&mockDish)
	assert.NoError(t, err)
	mockDishes := &models.DishesUseCase{}
	mockDishes.Dishes = append(mockDishes.Dishes, *mockDish)

	mockUCase.On("GetDishByRestaurant", mock.Anything, 1).Return(mockDishes)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "restaurant/" + slug, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	err = handler.GetDishesByRestaurants(c)
	require.NoError(t, err)
}
