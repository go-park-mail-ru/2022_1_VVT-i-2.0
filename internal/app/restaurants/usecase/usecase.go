package ucase

import (
	"strings"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/pkg/errors"
)

type RestaurantsUcase struct {
	RestaurantsRepo restaurants.Repository
}

func NewRestaurantsUcase(restaurantsRepo restaurants.Repository) *RestaurantsUcase {
	return &RestaurantsUcase{
		RestaurantsRepo: restaurantsRepo,
	}
}

func (u *RestaurantsUcase) GetAllRestaurants() (*models.RestaurantsUcase, error) {
	restaurantsRepoResp, err := u.RestaurantsRepo.GetRestaurants()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsResp := &models.RestaurantsUcase{Restaurants: make([]models.RestaurantUcase, len(restaurantsRepoResp.Restaurants))}

	for i, rest := range restaurantsRepoResp.Restaurants {
		restaurantsResp.Restaurants[i] = models.RestaurantUcase(rest)
	}

	return restaurantsResp, nil
}

func (u *RestaurantsUcase) GetRestaurantsByCategory(category models.GetRestaurantByCategoryUcaseReq) (*models.RestaurantsUcase, error) {
	restaurantsRepoResp, err := u.RestaurantsRepo.GetRestaurantsByCategory(models.GetRestaurantByCategoryRepoReq(category))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsResp := &models.RestaurantsUcase{Restaurants: make([]models.RestaurantUcase, len(restaurantsRepoResp.Restaurants))}

	for i, rest := range restaurantsRepoResp.Restaurants {
		restaurantsResp.Restaurants[i] = models.RestaurantUcase(rest)
	}

	return restaurantsResp, nil
}

func (u *RestaurantsUcase) GetRestaurantBySearchQuery(query models.GetRestaurantBySearchQueryUcaseReq) (*models.RestaurantsUcase, error) {
	query.Query = strings.Trim(query.Query, " \n\t")
	restaurantsRepoResp, err := u.RestaurantsRepo.GetRestaurantsBySearchQuery(models.GetRestaurantBySearchQueryRepoReq(query))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsResp := &models.RestaurantsUcase{Restaurants: make([]models.RestaurantUcase, len(restaurantsRepoResp.Restaurants))}

	for i, rest := range restaurantsRepoResp.Restaurants {
		restaurantsResp.Restaurants[i] = models.RestaurantUcase(rest)
	}

	return restaurantsResp, nil
}
