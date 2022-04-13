package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification"
)

type RestaurantsUsecase struct {
	Notificator 		notification.Notificator
	Cacher      		cacher.Cacher
	RestaurantsRepo    	restaurants.Repository
}

func NewRestaurantsUsecase(notificator notification.Notificator, cacher cacher.Cacher, restaurantsRepo restaurants.Repository) *RestaurantsUsecase {
	return &RestaurantsUsecase{
		Notificator: 		notificator,
		Cacher:      		cacher,
		RestaurantsRepo:    restaurantsRepo,
	}
}

func (u *RestaurantsUsecase) GetAllRestaurants() ([]*models.Restaurant, error) {
	restaurantsData, err := u.RestaurantsRepo.
}