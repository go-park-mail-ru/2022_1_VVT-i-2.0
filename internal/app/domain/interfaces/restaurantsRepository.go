package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type RestaurantsRepository struct {
	mock.Mock
}

func (r *RestaurantsRepository) GetRestaurants() ([]*models.RestaurantDataStorage, error) {
	restaurant := &models.RestaurantDataStorage{}
	restaurant = (*models.RestaurantDataStorage)(data.Rest)
	restaurantsDS := make([]*models.RestaurantDataStorage, 0)
	restaurantsDS = append(restaurantsDS, restaurant)
	return restaurantsDS, nil
}

func (r *RestaurantsRepository) GetRestaurantsBySlug(slug string) (*models.RestaurantDataStorage, error) {
	if slug == "" {
		return nil, nil
	}
	restaurant := &models.RestaurantDataStorage{}
	restaurant = (*models.RestaurantDataStorage)(data.Rest)
	return restaurant, nil
}

func (r *RestaurantsRepository) GetDishByRestaurants(id int) ([]*models.DishDataStorage, error) {
	if id == 0 {
		return nil, nil
	}
	dish := &models.DishDataStorage{}
	dish = (*models.DishDataStorage)(data.Dish)
	dishesDS := make([]*models.DishDataStorage, 0)
	dishesDS = append(dishesDS, dish)
	return dishesDS, nil
}

func (r *RestaurantsRepository) GetCommentsRestaurantByRestaurants(id int) ([]*models.CommentRestaurantDataStorage, error) {
	panic("implement me")
}

func (r *RestaurantsRepository) AddCommentsRestaurantByRestaurants(item *models.AddCommentRestaurantDataStorage) (*models.CommentRestaurantDataStorage, error) {
	panic("implement me")
}
