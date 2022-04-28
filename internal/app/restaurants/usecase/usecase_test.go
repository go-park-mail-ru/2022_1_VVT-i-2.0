package usecase

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)
import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"

func TestRestaurantsUsecase_GetAllRestaurants(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	restData, err := useCase.GetAllRestaurants()
	require.NoError(t, err)

	mockRestaurant := &models.RestaurantUsecase{}
	mockRestaurant = (*models.RestaurantUsecase)(data.Rest)

	if !reflect.DeepEqual(restData.Restaurants[0], mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], mockRestaurant)
		return
	}
}

func TestRestaurantsUsecase_GetRestaurantBySluf(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	restData, err := useCase.GetRestaurantBySluf("slug")
	require.NoError(t, err)

	mockRestaurant := &models.RestaurantUsecase{}
	mockRestaurant = (*models.RestaurantUsecase)(data.Rest)

	if !reflect.DeepEqual(restData, mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restData, mockRestaurant)
		return
	}
}

func TestRestaurantsUsecase_GetDishByRestaurant(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	dishData, err := useCase.GetDishByRestaurant(1)
	require.NoError(t, err)

	mockDish := &models.DishUseCase{}
	mockDish = (*models.DishUseCase)(data.Dish)

	if !reflect.DeepEqual(dishData.Dishes[0], mockDish) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
}

func TestRestaurantsUsecase_GetCommentsRestaurantByRestaurants(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	commentRestaurantData, err := useCase.GetCommentsRestaurantByRestaurants(1)
	require.NoError(t, err)

	mockDish := &models.CommentRestaurantUseCase{}
	mockDish = (*models.CommentRestaurantUseCase)(data.CommentRestaurant)

	if !reflect.DeepEqual(commentRestaurantData, mockDish) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData, data.Rest)
		return
	}
}
