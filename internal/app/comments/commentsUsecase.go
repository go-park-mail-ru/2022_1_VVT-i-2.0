package comments

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetRestaurantComments(slug string) (*models.CommentsRestaurantUseCase, error)
	AddRestaurantComment(id models.UserId, item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error)
	//AddRestaurantComment(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error)
}
