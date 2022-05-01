package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type RestaurantsUsecase struct {
	mock.Mock
}

func (a *CommentsUsecase) GetAllRestaurants() (*models.RestaurantsUcase, error) {
	mockRestaurant := (*models.RestaurantUcase)(data.Rest)
	mockRestaurants := &models.RestaurantsUcase{}
	mockRestaurants.Restaurants = append(mockRestaurants.Restaurants, *mockRestaurant)
	return mockRestaurants, nil
}
