package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type CommentsUsecase struct {
	mock.Mock
}

func (a *CommentsUsecase) GetRestaurantComments(id int) (*models.CommentsRestaurantUseCase, error) {
	if id != 0 {
		return nil, nil
	}
	mockCommentRestaurant := (*models.CommentRestaurantUseCase)(data.CommentRestaurant)
	mockmockCommentRestaurants := &models.CommentsRestaurantUseCase{}
	mockmockCommentRestaurants.Comment = append(mockmockCommentRestaurants.Comment, *mockCommentRestaurant)
	return mockmockCommentRestaurants, nil
}

func (a *CommentsUsecase) AddRestaurantComment(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	if item == nil {
		return nil, nil
	}
	mockCommentRestaurant := (*models.CommentRestaurantUseCase)(data.CommentRestaurant)
	return mockCommentRestaurant, nil
}
