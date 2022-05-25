package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type DishesRepo struct {
	mock.Mock
}

func (r *DishesRepo) GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.DishesRestaurantRepo, error) {
		if req.Slug == "" {
			return nil, nil
		}
		return &models.DishesRestaurantRepo{
			Id:                   1,
			Name:                 "Name",
			ImagePath:            "ImagePath",
			Slug:                 "Slug",
			MinPrice:             1,
			AggRating:            9,
			ReviewCount:          2,
			UpMinutesToDelivery:   3,
			DownMinutesToDelivery: 2}, nil
}

func (r *DishesRepo) GetCategories(req models.GetCategoriesByIdRepoReq) (*models.Categories, error) {
	if req.Id == 0 {
		return nil, nil
	}
	return &models.Categories{Categories: []string{"1"}}, nil
}

func (r *DishesRepo) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error) {
	if req.Id == 0 {
		return nil, nil
	}
	return &models.GetRestaurantDishesCategoriesRepoResp{Dishes: []models.DishCategoriesRepo{{Id: 1, Category: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"}}}, nil
}

type DishesRepoErr struct {
	mock.Mock
}

func (r *DishesRepoErr) GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.DishesRestaurantRepo, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (r *DishesRepoErr) GetCategories(req models.GetCategoriesByIdRepoReq) (*models.Categories, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (r *DishesRepoErr) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
