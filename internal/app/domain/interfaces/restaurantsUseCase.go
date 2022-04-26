package interfaces

import (
	"github.com/bxcodec/faker"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

type ArticleUsecase struct {
	mock.Mock
}

func (a ArticleUsecase) GetAllRestaurants() (*models.RestaurantsUsecase, error) {
	mockRestaurant := &models.RestaurantUsecase{}
	err := faker.FakeData(&mockRestaurant)
	if err != nil {
		return nil, errors.Wrap(err, "error")
	}
	mockRestaurants := &models.RestaurantsUsecase{}
	mockRestaurants.Restaurants = append(mockRestaurants.Restaurants, *mockRestaurant)
	return mockRestaurants, nil
}

func (a ArticleUsecase) GetRestaurantBySluf(slug string) (*models.RestaurantUsecase, error) {
	if slug != "" {
		return nil, nil
	}
	mockRestaurant := &models.RestaurantUsecase{}
	err := faker.FakeData(&mockRestaurant)
	if err != nil {
		return nil, errors.Wrap(err, "error")
	}
	return mockRestaurant, nil
}

func (a ArticleUsecase) GetDishByRestaurant(id int) (*models.DishesUseCase, error) {
	if id == 0 {
		return nil, nil
	}
	mockDish := &models.DishUseCase{}
	err := faker.FakeData(&mockDish)
	if err != nil {
		return nil, errors.Wrap(err, "error")
	}
	mockDishes := &models.DishesUseCase{}
	mockDishes.Dishes = append(mockDishes.Dishes, *mockDish)
	return mockDishes, nil
}

func (a ArticleUsecase) GetCommentsRestaurantByRestaurants(id int) (*models.CommentsRestaurantUseCase, error) {
	panic("implement me")
}

func (a ArticleUsecase) AddCommentsRestaurantByRestaurants(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	panic("implement me")
}