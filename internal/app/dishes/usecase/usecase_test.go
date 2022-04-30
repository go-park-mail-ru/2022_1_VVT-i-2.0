package usecase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/require"
)

func TestDishesUcase_GetRestaurantBySlug(t *testing.T) {
	mockDishesRepo := new(interfaces.DishesRepository)
	useCase := DishesUsecase{
		DishesRepo: mockDishesRepo,
	}

	restData, err := useCase.GetRestaurantBySlug("slug")
	require.NoError(t, err)

	mockRestaurant := (*models.RestaurantUcase)(data.Rest)

	if !reflect.DeepEqual(restData, mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restData, mockRestaurant)
		return
	}
}

func TestRestaurantsUsecase_GetDishByRestaurant(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.DishesRepository)
	useCase := DishesUsecase{
		DishesRepo: mockRestaurantsRepo,
	}

	dishData, err := useCase.GetDishesByRestaurant(1)
	require.NoError(t, err)

	mockDish := (*models.DishUcase)(data.Dish)

	if !reflect.DeepEqual(dishData.Dishes[0], *mockDish) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], mockDish)
		return
	}
}
