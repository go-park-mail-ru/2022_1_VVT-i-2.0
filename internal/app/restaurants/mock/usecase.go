package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type RestaurantsUcase struct {
	mock.Mock
}

func (a *RestaurantsUcase) GetAllRestaurants() (*models.RestaurantsUcase, error) {
	return &models.RestaurantsUcase{Restaurants: []models.RestaurantUcase{{
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

func (a *RestaurantsUcase) GetRestaurantsByCategory(category models.GetRestaurantByCategoryUcaseReq) (*models.RestaurantsUcase, error) {
	return &models.RestaurantsUcase{Restaurants: []models.RestaurantUcase{{
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

func (a *RestaurantsUcase) GetRestaurantBySearchQuery(query models.GetRestaurantBySearchQueryUcaseReq) (*models.RestaurantsUcase, error) {
	return &models.RestaurantsUcase{Restaurants: []models.RestaurantUcase{{
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

type RestaurantsUcaseErr struct {
	mock.Mock
}

func (a *RestaurantsUcaseErr) GetAllRestaurants() (*models.RestaurantsUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (a *RestaurantsUcaseErr) GetRestaurantsByCategory(category models.GetRestaurantByCategoryUcaseReq) (*models.RestaurantsUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (a *RestaurantsUcaseErr) GetRestaurantBySearchQuery(query models.GetRestaurantBySearchQueryUcaseReq) (*models.RestaurantsUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
