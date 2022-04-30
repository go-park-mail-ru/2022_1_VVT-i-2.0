package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type CommentsRepository struct {
	mock.Mock
}

func (r *CommentsRepository) GetRestaurantByID(id int) (*models.RestaurantDataStorage, error) {
	if id == 0 {
		return nil, nil
	}
	restaurant := (*models.RestaurantDataStorage)(data.Rest)
	return restaurant, nil
}

func (r *CommentsRepository) GetRestaurantComments(id int) ([]*models.CommentRestaurantDataStorage, error) {
	commentRestaurant := (*models.CommentRestaurantDataStorage)(data.CommentRestaurant)
	commentRestaurantsDS := make([]*models.CommentRestaurantDataStorage, 0)
	commentRestaurantsDS = append(commentRestaurantsDS, commentRestaurant)
	return commentRestaurantsDS, nil
}

func (r *CommentsRepository) AddRestaurantComment(item *models.AddCommentRestaurantDataStorage) (*models.CommentRestaurantDataStorage, error) {
	if item == nil {
		return nil, nil
	}
	commentRestaurant := (*models.CommentRestaurantDataStorage)(data.CommentRestaurant)
	return commentRestaurant, nil
}

func (r *CommentsRepository) UpdateRestaurantRating(restId int, newRestRating int, countRating int) (*models.RestaurantDataStorage, error) {
	if restId == 0 && newRestRating == 0 && countRating == 0 {
		return nil, nil
	}
	restaurant := (*models.RestaurantDataStorage)(data.Rest)
	return restaurant, nil
}
