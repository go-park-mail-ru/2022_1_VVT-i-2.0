package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurants() ([]*models.RestaurantDataStorage, error)
	GetRestaurantsByCategory(category models.GetRestaurantByCategoryRepoReq) ([]*models.RestaurantDataStorage, error)
	GetRestaurantsBySeachQuery(category models.GetRestaurantBySearchQueryRepoReq) ([]*models.RestaurantDataStorage, error)
}
