package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/pkg/errors"
)

type DishesUsecase struct {
	DishesRepo dishes.Repository
}

func NewDishesUsecase(restaurantsRepo dishes.Repository) *DishesUsecase {
	return &DishesUsecase{
		DishesRepo: restaurantsRepo,
	}
}

func (u *DishesUsecase) GetRestaurantBySlug(slug string) (*models.RestaurantUcase, error) {
	restaurantData, err := u.DishesRepo.GetRestaurantBySlug(slug)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}
	return &models.RestaurantUcase{
		Id:                   restaurantData.Id,
		Name:                 restaurantData.Name,
		ImagePath:            restaurantData.ImagePath,
		Slug:                 restaurantData.Slug,
		MinPrice:             restaurantData.MinPrice,
		AggRating:            restaurantData.AggRating,
		ReviewCount:          restaurantData.ReviewCount,
		UpMinutsToDelivery:   restaurantData.UpMinutsToDelivery,
		DownMinutsToDelivery: restaurantData.DownMinutsToDelivery,
	}, nil
}

func (u *DishesUsecase) GetDishesByRestaurant(id int) (*models.DishesUcase, error) {
	dishesData, err := u.DishesRepo.GetDishesByRestaurant(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	dishesUC := &models.DishesUcase{}

	for _, dish := range dishesData {
		item := &models.DishUcase{
			Id:           dish.Id,
			RestaurantId: dish.RestaurantId,
			Name:         dish.Name,
			Description:  dish.Description,
			ImagePath:    dish.ImagePath,
			Calories:     dish.Calories,
			Weight:       dish.Weight,
			Price:        dish.Price,
		}
		dishesUC.Dishes = append(dishesUC.Dishes, *item)
	}

	return dishesUC, nil
}
