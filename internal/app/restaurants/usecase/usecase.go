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

	restaurantsByNameRepoResp, err := u.RestaurantsRepo.GetRestaurantsByNameQuery(models.GetRestaurantBySearchQueryRepoReq(query))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsByCategoryRepoResp, err := u.RestaurantsRepo.GetRestaurantsByCategoryQuery(models.GetRestaurantBySearchQueryRepoReq(query))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsByDishRepoResp, err := u.RestaurantsRepo.GetRestaurantsByQueryDish(models.GetRestaurantBySearchQueryRepoReq(query))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsResp := &models.RestaurantsUcase{Restaurants: make([]models.RestaurantUcase, 0, len(restaurantsByNameRepoResp.Restaurants)+len(restaurantsByCategoryRepoResp.Restaurants)+len(restaurantsByDishRepoResp.Restaurants))}

	uniqueRestaurants := make(map[string]bool)
	for _, rest := range append(append(restaurantsByNameRepoResp.Restaurants, restaurantsByCategoryRepoResp.Restaurants...), restaurantsByDishRepoResp.Restaurants...) {
		if uniqueRestaurants[rest.Name] == false {
			restaurantsResp.Restaurants = append(restaurantsResp.Restaurants, models.RestaurantUcase(rest))
			uniqueRestaurants[rest.Name] = true
		}
	}

	return restaurantsResp, nil
}
