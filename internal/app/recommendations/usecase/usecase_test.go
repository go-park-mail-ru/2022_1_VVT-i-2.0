package usecase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/mock"
	"github.com/stretchr/testify/assert"
)

func TestRecommendationsUcase_GetRecommendations(t *testing.T) {
	mockRecommendationRepo := new(mock.RecommendationRepo)
	ucase := NewRecommendationsUcase(mockRecommendationRepo)

	recommendationData, err := ucase.GetRecommendations(&models.RecommendationsUcaseReq{RestId: 1, DishesId: []int64{1}})
	assert.NoError(t, err)

	expectResp := &models.RecommendationsUcaseResp{
		Dishes: []models.RecommendationUcase{
			{Id: 1, Category: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
			{Id: 2, Category: 2, Name: "Name2", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
		},
	}

	if !reflect.DeepEqual(recommendationData.Dishes, expectResp.Dishes) {
		t.Errorf("results not match,\n want %v,\n have %v", expectResp.Dishes, recommendationData.Dishes)
		return
	}
}

func TestRecommendationsUcase_GetRecommendations_Err(t *testing.T) {
	mockRecommendationRepo := new(mock.RecommendationRepoErr)
	ucase := NewRecommendationsUcase(mockRecommendationRepo)

	_, err := ucase.GetRecommendations(&models.RecommendationsUcaseReq{RestId: 1, DishesId: []int64{1}})
	assert.Error(t, err)
}
