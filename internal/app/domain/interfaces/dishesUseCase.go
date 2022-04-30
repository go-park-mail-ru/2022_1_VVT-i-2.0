package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type DishesUcase struct {
	mock.Mock
}

func (a *DishesUcase) GetRestaurantBySlug(slug string) (*models.RestaurantUcase, error) {
	if slug != "" {
		return nil, nil
	}
	mockRestaurant := (*models.RestaurantUcase)(data.Rest)
	return mockRestaurant, nil
}

func (a *DishesUcase) GetDishesByRestaurant(id int) (*models.DishesUcase, error) {
	if id == 0 {
		return nil, nil
	}
	mockDish := (*models.DishUcase)(data.Dish)
	mockDishes := &models.DishesUcase{}
	mockDishes.Dishes = append(mockDishes.Dishes, *mockDish)
	return mockDishes, nil
}
