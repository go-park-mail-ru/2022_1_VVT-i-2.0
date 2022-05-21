package usecase

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCommentsUsecase_GetRestaurantComment(t *testing.T) {
// 	mockCommentsRepo := new(mock.CommentsRepository)
// 	useCase := NewCommentsUsecase(mockCommentsRepo)

// 	commentRestaurantData, err := useCase.GetRestaurantComments("slug")
// 	assert.NoError(t, err)

// 	mockComment := &models.CommentsRestaurantUseCase{Comment: []models.CommentRestaurantUseCase{{
// 		RestaurantId: 1,
// 		Author:       "author",
// 		Text:         "text",
// 		Stars:        4,
// 		Date:         "date"}}}

// 	if !reflect.DeepEqual(commentRestaurantData, mockComment) {
// 		t.Errorf("results not match, want %v, have %v", commentRestaurantData.Comment[0], mockComment)
// 		return
// 	}
// }

// func TestCommentsUsecase_AddRestaurantComment(t *testing.T) {
// 	mockRestaurantsRepo := new(mock.CommentsRepository)
// 	useCase := CommentsUsecase{
// 		Repository: mockRestaurantsRepo,
// 	}

// 	id := 1

// 	item := &models.AddCommentRestaurantUseCase{
// 		Slug:          "slug",
// 		CommentText:   "text",
// 		CommentRating: 5,
// 	}
// 	commentRestaurantData, err := useCase.AddRestaurantComment(models.UserId(id), item)
// 	assert.NoError(t, err)

// 	mockComment := &models.CommentRestaurantUseCase{
// 		RestaurantId: 1,
// 		Author:       "author",
// 		Text:         "text",
// 		Stars:        4,
// 		Date:         "date"}

// 	if !reflect.DeepEqual(commentRestaurantData, mockComment) {
// 		t.Errorf("results not match, want %v, have %v", commentRestaurantData, mockComment)
// 		return
// 	}
// }
