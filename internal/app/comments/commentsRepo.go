package comments

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurantByID(id int) (*models.RestaurantDataStorage, error)
	GetRestaurantComments(id int) ([]*models.CommentRestaurantDataStorage, error)
	AddRestaurantComment(item *models.AddCommentRestaurantDataStorage) (*models.CommentRestaurantDataStorage, error)
	UpdateRestaurantRating(restId int, newRestRating int, countRating int) (*models.RestaurantDataStorage, error)
}
