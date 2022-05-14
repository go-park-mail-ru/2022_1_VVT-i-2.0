package dishes

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.DishesRestaurantRepo, error)
	GetCategories(req models.GetCategoriesByIdRepoReq) (*models.Categories, error)
	GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error)
}
