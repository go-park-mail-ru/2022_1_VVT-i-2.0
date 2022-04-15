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

func (u *RestaurantsUsecase) GetAllRestaurants() (*models.RestaurantsUsecase, error) {
	restaurantsData, err := u.RestaurantsRepo.GetRestaurants()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsUC := &models.RestaurantsUsecase{}

	for _, rest := range restaurantsData {
		item := &models.RestaurantUsecase{
			Id: rest.Id,
			Name: rest.Name,
			City: rest.City,
			Address: rest.Address,
			Image_path: rest.Image_path,
			Slug: rest.Slug,
			Min_price: rest.Min_price,
			Avg_price: rest.Avg_price,
			Rating: rest.Rating,
		}
		restaurantsUC.Restaurants = append(restaurantsUC.Restaurants, *item)
	}

	return restaurantsUC, nil
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

	dishesUC := &models.DishesDataStorage{}

	for _, dish := range dishesData {
		item := &models.DishDataStorage{
			Id: dish.Id,
			Restaurant: dish.Restaurant,
			Name: dish.Name,
			Description: dish.Description,
			Image_path: dish.Image_path,
			Calories: dish.Calories,
			Price: dish.Price,
		}
		dishesUC.Dishes = append(dishesUC.Dishes, *item)
	}

	return dishesUC, nil
}