package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurants() ([]*models.RestaurantDataStorage, error)
	GetRestaurantsBySlug(slug string) (*models.RestaurantDataStorage, error)
	GetDishByRestaurants(id int) ([]*models.DishDataStorage, error)
	GetCommentsRestaurantByRestaurants(id int) ([]*models.CommentRestaurantDataStorage, error)
	AddCommentsRestaurantByRestaurants(item *models.AddCommentRestaurantDataStorage) (bool, error)
}