package dishes

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetRestaurantBySlug(slug string) (*models.RestaurantUcase, error)
	GetDishesByRestaurant(id int) (*models.DishesUcase, error)
}
