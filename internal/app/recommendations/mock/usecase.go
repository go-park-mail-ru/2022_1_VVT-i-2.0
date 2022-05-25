package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type RecommendationsUcase struct {
	mock.Mock
}

func (u *RecommendationsUcase) GetRecommendations(req models.RecommendationsOrderListsUsecaseReq) (*models.DishRecommendationListsUsecase, error) {
	return &models.DishRecommendationListsUsecase{
		Dishes: []models.DishRecommendationUsecase{
			{Id: 2, Category: 2, Name: "Name2", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
			{Id: 3, Category: 3, Name: "Name3", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
		},
	}, nil
}

type RecommendationsUcaseErr struct {
	mock.Mock
}

func (u *RecommendationsUcaseErr) GetRecommendations(req models.RecommendationsOrderListsUsecaseReq) (*models.DishRecommendationListsUsecase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
