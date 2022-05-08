package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type DishesUcase struct {
	mock.Mock
}

func (u *DishesUcase) GetRestaurantDishes(req models.GetRestaurantDishesUcaseReq) (*models.GetRestaurantDishesUcaseResp, error) {
	if req.Slug == "" {
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
	}

	return &models.GetRestaurantDishesUcaseResp{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "Slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2,
		Dishes:               []models.DishUcase{{Id: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"}},
	}, nil
}

type DishesUcaseErr struct {
	mock.Mock
}

func (a *DishesUcaseErr) GetRestaurantDishes(req models.GetRestaurantDishesUcaseReq) (*models.GetRestaurantDishesUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.UNKNOWN_ERROR, "")
}
