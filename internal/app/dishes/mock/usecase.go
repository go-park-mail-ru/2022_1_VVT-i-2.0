package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type DishesUcase struct {
	mock.Mock
}

func (u *DishesUcase) GetRestaurantDishes(req models.GetRestaurantDishesUcaseReq) (*models.GetRestaurantDishesCategoriesUcaseResp, error) {
	if req.Slug == "" {
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
	}

	return &models.GetRestaurantDishesCategoriesUcaseResp{
		Id:                  	 1,
		Name:                 	"Name",
		ImagePath:            	"ImagePath",
		Slug:                 	"Slug",
		MinPrice:             	1,
		AggRating:            	9,
		ReviewCount:          	2,
		UpMinutesToDelivery:   	3,
		DownMinutesToDelivery: 	2,
		Dishes:               	[]models.DishCategoriesUsecase{{Id: 1, Category: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"}},
		Categories: 			[]models.CategoriesDishesUcaseResp{{Categories: "1", Dishes: []int{1}}},
	}, nil
}

type DishesUcaseErr struct {
	mock.Mock
}

func (a *DishesUcaseErr) GetRestaurantDishes(req models.GetRestaurantDishesUcaseReq) (*models.GetRestaurantDishesCategoriesUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.UNKNOWN_ERROR, "")
}
