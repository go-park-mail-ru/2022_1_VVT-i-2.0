package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type RestaurantsUsecase struct {
	mock.Mock
}

func (a *RestaurantsUsecase) GetAllRestaurants() (*models.RestaurantsUsecase, error) {
	mockRestaurant := &models.RestaurantUsecase{}
	mockRestaurant = (*models.RestaurantUsecase)(data.Rest)
	mockRestaurants := &models.RestaurantsUsecase{}
	mockRestaurants.Restaurants = append(mockRestaurants.Restaurants, *mockRestaurant)
	return mockRestaurants, nil
}

func (a *RestaurantsUsecase) GetRestaurantBySluf(slug string) (*models.RestaurantUsecase, error) {
	if slug != "" {
		return nil, nil
	}
	mockRestaurant := &models.RestaurantUsecase{}
	mockRestaurant = (*models.RestaurantUsecase)(data.Rest)
	return mockRestaurant, nil
}

func (a *RestaurantsUsecase) GetDishByRestaurant(id int) (*models.DishesUseCase, error) {
	if id == 0 {
		return nil, nil
	}
	mockDish := &models.DishUseCase{}
	mockDish = (*models.DishUseCase)(data.Dish)
	mockDishes := &models.DishesUseCase{}
	mockDishes.Dishes = append(mockDishes.Dishes, *mockDish)
	return mockDishes, nil
}

func (a *RestaurantsUsecase) GetCommentsRestaurantByRestaurants(id int) (*models.CommentsRestaurantUseCase, error) {
	if id != 0 {
		return nil, nil
	}
	mockCommentRestaurant := &models.CommentRestaurantUseCase{}
	mockCommentRestaurant = (*models.CommentRestaurantUseCase)(data.CommentRestaurant)
	mockmockCommentRestaurants := &models.CommentsRestaurantUseCase{}
	mockmockCommentRestaurants.Comment = append(mockmockCommentRestaurants.Comment, *mockCommentRestaurant)
	return mockmockCommentRestaurants, nil
}

func (a *RestaurantsUsecase) AddCommentsRestaurantByRestaurants(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	if item == nil {
		return nil, nil
	}
	mockCommentRestaurant := &models.CommentRestaurantUseCase{}
	mockCommentRestaurant = (*models.CommentRestaurantUseCase)(data.CommentRestaurant)
	return mockCommentRestaurant, nil
}