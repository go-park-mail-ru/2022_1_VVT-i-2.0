package usecase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/require"
)

func TestCommentsUsecase_GetRestaurantComment(t *testing.T) {
	mockCommentsRepo := new(interfaces.CommentsRepository)
	useCase := CommentsUsecase{
		Repository: mockCommentsRepo,
	}

	commentRestaurantData, err := useCase.GetRestaurantComments(1)
	require.NoError(t, err)

	mockComment := (*models.CommentRestaurantUseCase)(data.CommentRestaurant)

	if !reflect.DeepEqual(commentRestaurantData.Comment[0], *mockComment) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData.Comment[0], mockComment)
		return
	}
}

func TestCommentsUsecase_AddRestaurantComment(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.CommentsRepository)
	useCase := CommentsUsecase{
		Repository: mockRestaurantsRepo,
	}

	item := &models.AddCommentRestaurantUseCase{
		Restaurant:     1,
		User_id:        1,
		Comment_text:   "text",
		Comment_rating: 5,
	}
	commentRestaurantData, err := useCase.AddRestaurantComment(item)
	require.NoError(t, err)

	mockComment := (*models.CommentRestaurantUseCase)(data.CommentRestaurant)

	if !reflect.DeepEqual(commentRestaurantData, mockComment) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData, mockComment)
		return
	}
}
