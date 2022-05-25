package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type CommentsUsecase struct {
	mock.Mock
}

func (a *CommentsUsecase) GetRestaurantComments(req models.GetRestaurantCommentsUcaseReq) (*models.CommentsRestaurantUseCase, error) {
	if req.Slug == "" {
		return nil, nil
	}
	return &models.CommentsRestaurantUseCase{Comment: []models.CommentRestaurantUseCase{{
		RestaurantId: 1,
		Author:       "author",
		Text:         "text",
		Stars:        4,
		Date:         "date"}}}, nil
}

func (a *CommentsUsecase) AddRestaurantComment(req models.AddCommentRestaurantUcaseReq) (*models.CommentRestaurantUseCase, error) {
	return &models.CommentRestaurantUseCase{
		RestaurantId: 1,
		Author:       "author",
		Text:         "text",
		Stars:        4,
		Date:         "date"}, nil
}

type CommentsUsecaseErr struct {
	mock.Mock
}

func (a *CommentsUsecaseErr) GetRestaurantComments(req models.GetRestaurantCommentsUcaseReq) (*models.CommentsRestaurantUseCase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (a *CommentsUsecaseErr) AddRestaurantComment(req models.AddCommentRestaurantUcaseReq) (*models.CommentRestaurantUseCase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

