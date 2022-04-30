package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type DishesRepository struct {
	mock.Mock
}

func (r *DishesRepository) GetRestaurantBySlug(slug string) (*models.RestaurantDataStorage, error) {
	if slug == "" {
		return nil, nil
	}
	restaurant := (*models.RestaurantDataStorage)(data.Rest)
	return restaurant, nil
}

func (r *DishesRepository) GetRestaurantByID(id int) (*models.RestaurantDataStorage, error) {
	if id == 0 {
		return nil, nil
	}
	restaurant := (*models.RestaurantDataStorage)(data.Rest)
	return restaurant, nil
}

func (r *DishesRepository) GetDishesByRestaurant(id int) ([]*models.DishDataStorage, error) {
	if id == 0 {
		return nil, nil
	}
	dish := (*models.DishDataStorage)(data.Dish)
	dishesDS := make([]*models.DishDataStorage, 0)
	dishesDS = append(dishesDS, dish)
	return dishesDS, nil
}
