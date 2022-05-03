package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetAllRestaurants() (*models.RestaurantsUcase, error)
	GetRestaurantsByCategory(category models.GetRestaurantByCategoryUcaseReq) (*models.RestaurantsUcase, error)
	GetRestaurantBySearchQuery(category models.GetRestaurantBySearchQueryUcaseReq) (*models.RestaurantsUcase, error)
}
