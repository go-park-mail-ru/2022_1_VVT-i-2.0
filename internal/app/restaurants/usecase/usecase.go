package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/pkg/errors"
)

type RestaurantsUsecase struct {
	RestaurantsRepo    	restaurants.Repository
}

func NewRestaurantsUsecase(restaurantsRepo restaurants.Repository) *RestaurantsUsecase {
	return &RestaurantsUsecase{
		RestaurantsRepo:    restaurantsRepo,
	}
}

func (u *RestaurantsUsecase) GetAllRestaurants() (*models.RestaurantsDataStorage, error) {
	restaurantsData, err := u.RestaurantsRepo.GetRestaurants()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}
	return &models.RestaurantsDataStorage{
		Restaurants: restaurantsData.Restaurants,
	}, nil
}

func (u *RestaurantsUsecase) GetRestaurantBySluf(slug string) (*models.RestaurantUsecase, error) {
	restaurantData, err := u.RestaurantsRepo.GetRestaurantsBySlug(slug)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}
	return &models.RestaurantUsecase{
		Id: restaurantData.Id,
		Name: restaurantData.Name,
		City: restaurantData.City,
		Address: restaurantData.Address,
		Image_path: restaurantData.Image_path,
		Slug: restaurantData.Slug,
		Min_price: restaurantData.Min_price,
		Avg_price: restaurantData.Avg_price,
		Rating: restaurantData.Rating,
	}, nil
}

func (u *RestaurantsUsecase) GetDishByRestaurant(id int) (*models.DishesDataStorage, error) {
	dishesData, err := u.RestaurantsRepo.GetDishByRestaurants(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}
	return &models.DishesDataStorage{
		Dishes: dishesData,
	}, nil
}