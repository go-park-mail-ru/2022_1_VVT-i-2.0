package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Restaurants interface {
	GetRestaurants() ([]*models.Restaurant, error)
	GetRestaurantsBySlug(slug string) (*models.Restaurant, error)
	GetDishByRestaurants(id int) ([]*models.Dish, error)
	//GetDishByRestaurants(slug string) (*models.Restaurant, *models.Dish, error)
}