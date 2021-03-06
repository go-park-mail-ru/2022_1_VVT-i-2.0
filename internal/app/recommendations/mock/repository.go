package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type RecommendationRepo struct {
	mock.Mock
}

func (r *RecommendationRepo) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error) {
	if req.Id == 0 {
		return nil, nil
	}
	return &models.GetRestaurantDishesCategoriesRepoResp{Dishes: []models.DishCategoriesRepo{
		{Id: 1, Category: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
		{Id: 2, Category: 2, Name: "Name2", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"},
	}}, nil
}

type RecommendationRepoErr struct {
	mock.Mock
}

func (r *RecommendationRepoErr) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}