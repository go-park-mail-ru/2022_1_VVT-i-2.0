package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurants() (*models.RestaurantsRepo, error)
	GetRestaurantsByCategory(category models.GetRestaurantByCategoryRepoReq) (*models.RestaurantsRepo, error)
	GetRestaurantsBySeachQuery(category models.GetRestaurantBySearchQueryRepoReq) (*models.RestaurantsRepo, error)
}
