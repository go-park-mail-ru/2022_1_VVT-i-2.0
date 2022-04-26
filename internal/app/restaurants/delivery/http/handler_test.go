package restaurantsHandler

import (
	"github.com/bxcodec/faker"
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ArticleUsecase struct {
	mock.Mock
}

func (a ArticleUsecase) GetAllRestaurants() (*models.RestaurantsUsecase, error) {
	mockRestaurant := &models.RestaurantUsecase{}
	err := faker.FakeData(&mockRestaurant)
	if err != nil {
		return nil, errors.Wrap(err, "error")
	}
	mockRestaurants := &models.RestaurantsUsecase{}
	mockRestaurants.Restaurants = append(mockRestaurants.Restaurants, *mockRestaurant)
	return mockRestaurants, nil
}

func (a ArticleUsecase) GetRestaurantBySluf(slug string) (*models.RestaurantUsecase, error) {
	panic("implement me")
}

func (a ArticleUsecase) GetDishByRestaurant(id int) (*models.DishesUseCase, error) {
	panic("implement me")
}

func (a ArticleUsecase) GetCommentsRestaurantByRestaurants(id int) (*models.CommentsRestaurantUseCase, error) {
	panic("implement me")
}

func (a ArticleUsecase) AddCommentsRestaurantByRestaurants(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	panic("implement me")
}

func TestRestaurantsHandler_GetAllRestaurants(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockRestaurant := &models.RestaurantUsecase{}
	err = faker.FakeData(&mockRestaurant)
	assert.NoError(t, err)
	mockRestaurants := &models.RestaurantsUsecase{}
	mockRestaurants.Restaurants = append(mockRestaurants.Restaurants, *mockRestaurant)

	mockUCase := new(ArticleUsecase)
	mockUCase.On("Fetch", mock.Anything).Return(mockRestaurants)

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
