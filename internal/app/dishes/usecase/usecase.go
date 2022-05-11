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

	categories, err := u.Repo.GetCategories(restaurant.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant categories")
	}

	dishes, err := u.Repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: restaurant.Id})
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant dishes")
	}

	Resp := &models.GetRestaurantDishesCategoriesUcaseResp{
			Id:                   restaurant.Id,
			Name:                 restaurant.Name,
			ImagePath:            restaurant.ImagePath,
			Slug:                 restaurant.Slug,
			MinPrice:             restaurant.MinPrice,
			AggRating:            restaurant.AggRating,
			ReviewCount:          restaurant.ReviewCount,
			UpMinutsToDelivery:   restaurant.UpMinutsToDelivery,
			DownMinutsToDelivery: restaurant.DownMinutsToDelivery,
	}

	for _, item := range categories.Categories {
		catDis := models.CategoriesDishes{
			Categories: item,
		}
		Resp.Dishes = append(Resp.Dishes, catDis)
	}

	for _, item := range dishes.Dishes {
		var car = item.Categori
		Resp.Dishes[car-1].Dishes = append(Resp.Dishes[car-1].Dishes, item)
	}

	return Resp, nil
}

//func (u *DishesUcase) GetRestaurantDishes(req models.GetRestaurantDishesUcaseReq) (*models.GetRestaurantDishesUcaseResp, error) {
//	restaurant, err := u.Repo.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq(req))
//	if err != nil {
//		return nil, errors.Wrap(err, "error getting restaurant")
//	}
//	dishes, err := u.Repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: restaurant.Id})
//	if err != nil {
//		return nil, errors.Wrap(err, "error getting restaurant dishes")
//	}
//
//	Resp := &models.GetRestaurantDishesUcaseResp{
//		Id:                   restaurant.Id,
//		Name:                 restaurant.Name,
//		ImagePath:            restaurant.ImagePath,
//		Slug:                 restaurant.Slug,
//		MinPrice:             restaurant.MinPrice,
//		AggRating:            restaurant.AggRating,
//		ReviewCount:          restaurant.ReviewCount,
//		UpMinutsToDelivery:   restaurant.UpMinutsToDelivery,
//		DownMinutsToDelivery: restaurant.DownMinutsToDelivery,
//		Dishes:               make([]models.DishUcase, len(dishes.Dishes)),
//	}
//
//	for i, dish := range dishes.Dishes {
//		Resp.Dishes[i] = models.DishUcase(dish)
//	}
//
//	return Resp, nil
//}
