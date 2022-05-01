package dishes

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurantBySlug(slug string) (*models.RestaurantDataStorage, error)
	GetDishesByRestaurant(id int) ([]*models.DishDataStorage, error)
}
