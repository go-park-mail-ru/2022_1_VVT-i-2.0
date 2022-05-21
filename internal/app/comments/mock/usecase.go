package mock

// import (
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
// 	"github.com/stretchr/testify/mock"
// )

// type CommentsUsecase struct {
// 	mock.Mock
// }

// func (a *CommentsUsecase) GetRestaurantComments(slug string) (*models.CommentsRestaurantUseCase, error) {
// 	if slug != "" {
// 		return nil, nil
// 	}
// 	return &models.CommentsRestaurantUseCase{Comment: []models.CommentRestaurantUseCase{{
// 		RestaurantId: 1,
// 		Author:       "author",
// 		Text:         "text",
// 		Stars:        4,
// 		Date:         "date"}}}, nil
// }

// func (a *CommentsUsecase) AddRestaurantComment(id models.UserId, item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
// 	if item == nil {
// 		return nil, nil
// 	}
// 	return &models.CommentRestaurantUseCase{
// 		RestaurantId: 1,
// 		Author:       "author",
// 		Text:         "text",
// 		Stars:        4,
// 		Date:         "date"}, nil
// }

// type CommentsUsecaseErr struct {
// 	mock.Mock
// }

// func (a *CommentsUsecaseErr) GetRestaurantComments(slug string) (*models.CommentsRestaurantUseCase, error) {
// 	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
// }

// func (a *CommentsUsecaseErr) AddRestaurantComment(id models.UserId, item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
// 	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
// }
