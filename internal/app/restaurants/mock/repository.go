package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type RestaurantsRepo struct {
	mock.Mock
}

func (r *RestaurantsRepo) GetRestaurants() (*models.RestaurantsRepo, error) {
	return &models.RestaurantsRepo{Restaurants: []models.RestaurantRepo{{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}}}, nil
}

func (r *RestaurantsRepo) GetRestaurantsByCategory(category models.GetRestaurantByCategoryRepoReq) (*models.RestaurantsRepo, error) {
	return &models.RestaurantsRepo{Restaurants: []models.RestaurantRepo{{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}}}, nil
}

func (r *RestaurantsRepo) GetRestaurantsBySearchQuery(category models.GetRestaurantBySearchQueryRepoReq) (*models.RestaurantsRepo, error) {
	return &models.RestaurantsRepo{Restaurants: []models.RestaurantRepo{{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}}}, nil
}
