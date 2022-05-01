package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/pkg/errors"
)

type RestaurantsUsecase struct {
	RestaurantsRepo restaurants.Repository
}

func NewRestaurantsUsecase(restaurantsRepo restaurants.Repository) *RestaurantsUsecase {
	return &RestaurantsUsecase{
		RestaurantsRepo: restaurantsRepo,
	}
}

func (u *RestaurantsUsecase) GetAllRestaurants() (*models.RestaurantsUcase, error) {
	restaurantsData, err := u.RestaurantsRepo.GetRestaurants()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurants := &models.RestaurantsUcase{}

	for _, rest := range restaurantsData {
		item := models.RestaurantUcase{
			Id:                   rest.Id,
			Name:                 rest.Name,
			ImagePath:            rest.ImagePath,
			Slug:                 rest.Slug,
			MinPrice:             rest.MinPrice,
			AggRating:            rest.AggRating,
			ReviewCount:          rest.ReviewCount,
			UpMinutsToDelivery:   rest.UpMinutsToDelivery,
			DownMinutsToDelivery: rest.DownMinutsToDelivery,
		}
		restaurants.Restaurants = append(restaurants.Restaurants, item)
	}

	return restaurants, nil
}
