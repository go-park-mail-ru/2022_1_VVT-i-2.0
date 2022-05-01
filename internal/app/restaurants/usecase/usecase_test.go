package usecase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/require"
)

func TestRestaurantsUsecase_GetAllRestaurants(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.DishesRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockRestaurantsRepo,
	}

	restData, err := useCase.GetAllRestaurants()
	require.NoError(t, err)

	mockRestaurant := (*models.RestaurantUcase)(data.Rest)

	if !reflect.DeepEqual(restData.Restaurants[0], *mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], mockRestaurant)
		return
	}
}
