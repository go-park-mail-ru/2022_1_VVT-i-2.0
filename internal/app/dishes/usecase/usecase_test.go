package ucase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestDishesUcase_RestaurantDishes(t *testing.T) {
	mockRestaurantsRepo := new(mock.DishesRepo)
	ucase := NewDishesUcase(mockRestaurantsRepo)

	dishData, err := ucase.GetRestaurantDishes(models.GetRestaurantDishesUcaseReq{Slug: "slug"})
	assert.NoError(t, err)

	expectResp := &models.GetRestaurantDishesCategoriesUcaseResp{
		Id:                   	1,
		Name:                 	"Name",
		ImagePath:            	"ImagePath",
		Slug:                 	"Slug",
		MinPrice:             	1,
		AggRating:            	9,
		ReviewCount:         	2,
		UpMinutesToDelivery:   	3,
		DownMinutesToDelivery:	2,
		Dishes:               	[]models.DishCategoriesUsecase{{Id: 1, Category: 1, Name: "Name", Description: "Description", Price: 10, Weight: 100, Calories: 200, RestaurantId: 1, ImagePath: "DishImagePath"}},
		Categories:				[]models.CategoriesDishesUcaseResp{{Categories: "1", Dishes: []int{1}}},
	}

	if !reflect.DeepEqual(dishData, expectResp) {
		t.Errorf("results not match,\n want %v,\n have %v", dishData.Dishes[0], expectResp)
		return
	}
}

func TestDishesUcase_RestaurantDishes_EmptySlug(t *testing.T) {
	mockRestaurantsRepo := new(mock.DishesRepoErr)
	ucase := NewDishesUcase(mockRestaurantsRepo)

	dishData, _ := ucase.GetRestaurantDishes(models.GetRestaurantDishesUcaseReq{Slug: "slug"})

	if reflect.DeepEqual(dishData, nil) {
		t.Errorf("results not match, want %v, have %v", dishData, nil)
		return
	}
}
