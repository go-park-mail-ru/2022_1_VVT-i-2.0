package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurants() (*models.RestaurantsDataStorage, error)
	GetRestaurantsBySlug(slug string) (*models.RestaurantDataStorage, error)
	GetDishByRestaurants(id int) ([]*models.Dish, error)
	//GetDishByRestaurants(slug string) (*models.Restaurant, *models.Dish, error)
}