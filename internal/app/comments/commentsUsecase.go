package comments

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetRestaurantComments(req models.GetRestaurantCommentsUcaseReq) (*models.CommentsRestaurantUseCase, error)
	AddRestaurantComment(req models.AddCommentRestaurantUcaseReq) (*models.CommentRestaurantUseCase, error)
}
