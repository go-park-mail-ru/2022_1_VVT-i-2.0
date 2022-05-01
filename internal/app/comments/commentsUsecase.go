package comments

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetRestaurantComments(id int) (*models.CommentsRestaurantUseCase, error)
	AddRestaurantComment(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error)
}
