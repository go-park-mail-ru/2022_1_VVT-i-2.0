package dishes

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantRepo, error)

	//GetCategories(id int) (*[]string, error)
	GetCategories(id int) (*models.Categories, error)

	//GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantCategoriesRepo, error)
	//GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesRepoResp, error)
	GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error)
}
