package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetAllRestaurants() (*models.RestaurantsUsecase, error)
	GetRestaurantBySluf(slug string) (*models.RestaurantUsecase, error)
	GetDishByRestaurant(id int) (*models.DishesDataStorage, error)
}
