package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations"
	"github.com/pkg/errors"
	"math/rand"
)

type RecommendationsUcase struct {
	Repo recommendations.Repository
}

func NewRecommendationsUcase(recommendationsRepo recommendations.Repository) *RecommendationsUcase {
	return &RecommendationsUcase{
		Repo: recommendationsRepo,
	}
}

func (u *RecommendationsUcase) GetRecommendations(req models.RecommendationsOrderListsUsecaseReq) (*models.DishRecommendationListsUsecase, error) {
	dishes, err := u.Repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: models.Id(req.RestId)})
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant dishes")
	}

	var categories []int
	for _, item := range dishes.Dishes {
		for _, item2 := range req.DishesId {
			if int64(item.Id) == item2 {
				var che = false
				for _, item3 := range categories{
					if item.Category ==item3 {
						che = true
					}
				}
				if !che {
					categories = append(categories, item.Category)
				}
			}
		}
	}

	recommendations := make([]models.DishCategoriesRepo, 0)

	for _, item := range dishes.Dishes {
		var che = false
		for _, item1 := range categories {
			if item.Category == item1 {
				che = true
			}
		}
		if !che {
			recommendations = append(recommendations, item)
		}
	}

	if len(recommendations) == 0 {
		fmt.Println("aue")
		return nil, nil
	}

	finalRecommendations := &models.DishRecommendationListsUsecase{
		Dishes: make([]models.DishRecommendationUsecase, 2),
	}

	for i := range finalRecommendations.Dishes{
		var che = rand.Intn(len(recommendations))
		var itme = models.DishRecommendationUsecase{
			Id:				recommendations[che].Id,
			Category:		recommendations[che].Category,
			RestaurantId:	recommendations[che].RestaurantId,
			Name:			recommendations[che].Name,
			Description:	recommendations[che].Description,
			ImagePath:		recommendations[che].ImagePath,
			Calories:		recommendations[che].Calories,
			Price:			recommendations[che].Price,
			Weight:			recommendations[che].Weight,
		}
		finalRecommendations.Dishes[i] = itme
	}

	return finalRecommendations, nil
}
