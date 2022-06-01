package usecase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews/mock"
	"github.com/stretchr/testify/assert"
)

func TestReviewsUcase_GetRestaurantReviews(t *testing.T) {
	mockCommentsRepo := new(mock.CommentsRepository)
	useCase := NewRestaurantReviewsUcase(mockCommentsRepo)

	commentRestaurantData, err := useCase.GetRestaurantReviews(&models.GetRestaurantReviewsUcaseReq{Slug: "slug"})
	assert.NoError(t, err)

	mockComment := &models.GetRestaurantReviewsUcaseResp{Reviews: []models.RestaurantReviewUcase{{
		Author: "author",
		Text:   "text",
		Stars:  4,
		Date:   "date"}}}

	if !reflect.DeepEqual(commentRestaurantData, mockComment) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData.Reviews[0], mockComment)
		return
	}
}

func TestCommentsUsecase_AddRestaurantComment(t *testing.T) {
	mockCommentsRepo := new(mock.CommentsRepository)
	useCase := RestaurantReviewsUcase{
		Repo: mockCommentsRepo,
	}

	commentRestaurantData, err := useCase.AddRestaurantReview(&models.AddRestaurantReviewUcaseReq{
		UserId: 1,
		Slug:   "slug",
		Text:   "text",
		Rating: 5,
	})
	assert.NoError(t, err)

	mockComment := &models.AddRestaurantReviewUcaseResp{
		Author: "author",
		Text:   "text",
		Stars:  4,
		Date:   "date"}

	if !reflect.DeepEqual(commentRestaurantData, mockComment) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData, mockComment)
		return
	}
}

func TestCommentsUsecase_GetRestaurantComment_Err(t *testing.T) {
	mockCommentsRepo := new(mock.CommentsRepositoryErr)
	useCase := NewRestaurantReviewsUcase(mockCommentsRepo)

	_, err := useCase.GetRestaurantReviews(&models.GetRestaurantReviewsUcaseReq{Slug: "slug"})
	assert.Error(t, err)
}

func TestCommentsUsecase_AddRestaurantComment_Err(t *testing.T) {
	mockCommentsRepo := new(mock.CommentsRepositoryErr)
	useCase := NewRestaurantReviewsUcase(mockCommentsRepo)

	_, err := useCase.AddRestaurantReview(&models.AddRestaurantReviewUcaseReq{
		UserId: 1,
		Slug:   "slug",
		Text:   "text",
		Rating: 5,
	})

	assert.Error(t, err)
}
