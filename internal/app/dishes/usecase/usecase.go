package ucase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/pkg/errors"
)

type DishesUcase struct {
	Repo dishes.Repository
}

func NewDishesUcase(restaurantsRepo dishes.Repository) *DishesUcase {
	return &DishesUcase{
		Repo: restaurantsRepo,
	}
}

func (u *DishesUcase) GetRestaurantDishes(req models.GetRestaurantDishesUcaseReq) (*models.GetRestaurantDishesCategoriesUcaseResp, error) {
	restaurant, err := u.Repo.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq(req))
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant")
	}

	categories, err := u.Repo.GetCategories(models.GetCategoriesByIdRepoReq{Id: models.Id(restaurant.Id)})
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant categories")
	}

	dishes, err := u.Repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: models.Id(restaurant.Id)})
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant dishes")
	}

	Resp := &models.GetRestaurantDishesCategoriesUcaseResp{
		Id:                    restaurant.Id,
		Name:                  restaurant.Name,
		ImagePath:             restaurant.ImagePath,
		Slug:                  restaurant.Slug,
		MinPrice:              restaurant.MinPrice,
		AggRating:             restaurant.AggRating,
		ReviewCount:           restaurant.ReviewCount,
		UpMinutesToDelivery:   restaurant.UpMinutesToDelivery,
		DownMinutesToDelivery: restaurant.DownMinutesToDelivery,
		Dishes:                make([]models.DishCategoriesUsecase, len(dishes.Dishes)),
		Categories:            make([]models.CategoriesDishesUcaseResp, len(categories.Categories)),
	}

	for i, dish := range dishes.Dishes {
		Resp.Dishes[i] = models.DishCategoriesUsecase(dish)
	}

	for i, item := range categories.Categories {
		Resp.Categories[i].Categories = item
	}

	for _, item := range dishes.Dishes {
		Resp.Categories[item.Category-1].Dishes = append(Resp.Categories[item.Category-1].Dishes, item.Id)
	}

	return Resp, nil
}
