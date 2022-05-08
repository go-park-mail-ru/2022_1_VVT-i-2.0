package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type DishesRepo struct {
	mock.Mock
}

func (r *DishesRepo) GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantRepo, error) {
	if req.Slug == "" {
		return nil, nil
	}
	return &models.RestaurantRepo{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "Slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}, nil
}

func (r *DishesRepo) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesRepoResp, error) {
	if req.Id == 0 {
		return nil, nil
	}
	return &models.GetRestaurantDishesRepoResp{Dishes: []models.DishRepo{{Id: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"}}}, nil
}
