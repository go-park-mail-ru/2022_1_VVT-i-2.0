package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/mock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRecommendationsUcase_GetRecommendations(t *testing.T) {
	mockRecommendationRepo := new(mock.RecommendationRepo)
	ucase := NewRecommendationsUcase(mockRecommendationRepo)

	recommendationData, err := ucase.GetRecommendations(models.RecommendationsOrderListsUsecaseReq{RestId: 1, DishesId: []int64{1}})
	assert.NoError(t, err)

	expectResp := &models.DishRecommendationListsUsecase{
		Dishes: []models.DishRecommendationUsecase{
			{Id: 2, Category: 2, Name: "Name2", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
			{Id: 2, Category: 2, Name: "Name2", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
		},
	}

	if !reflect.DeepEqual(recommendationData.Dishes, expectResp.Dishes) {
		t.Errorf("results not match,\n want %v,\n have %v", recommendationData.Dishes, expectResp.Dishes)
		return
	}
}