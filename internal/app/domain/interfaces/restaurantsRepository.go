package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type RestaurantsRepository struct {
	mock.Mock
}

func (r *DishesRepository) GetRestaurants() ([]*models.RestaurantDataStorage, error) {
	restaurant := (*models.RestaurantDataStorage)(data.Rest)
	restaurantsDS := make([]*models.RestaurantDataStorage, 0)
	restaurantsDS = append(restaurantsDS, restaurant)
	return restaurantsDS, nil
}
