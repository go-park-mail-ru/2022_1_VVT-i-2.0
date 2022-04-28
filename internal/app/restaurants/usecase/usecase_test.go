package usecase

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)
import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"

func TestRestaurantsUsecase_GetAllRestaurants(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockRestaurantsRepo,
	}

	restData, err := useCase.GetAllRestaurants()
	require.NoError(t, err)

	mockRestaurant := &models.RestaurantUsecase{}
	mockRestaurant = (*models.RestaurantUsecase)(data.Rest)

	if !reflect.DeepEqual(restData.Restaurants[0], *mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], mockRestaurant)
		return
	}
}

func TestRestaurantsUsecase_GetRestaurantBySluf(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockRestaurantsRepo,
	}

	restData, err := useCase.GetRestaurantBySluf("slug")
	require.NoError(t, err)

	mockRestaurant := &models.RestaurantUsecase{}
	mockRestaurant = (*models.RestaurantUsecase)(data.Rest)

	if !reflect.DeepEqual(restData, mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restData, mockRestaurant)
		return
	}
}

func TestRestaurantsUsecase_GetDishByRestaurant(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockRestaurantsRepo,
	}

	dishData, err := useCase.GetDishByRestaurant(1)
	require.NoError(t, err)

	mockDish := &models.DishUseCase{}
	mockDish = (*models.DishUseCase)(data.Dish)

	if !reflect.DeepEqual(dishData.Dishes[0], *mockDish) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], mockDish)
		return
	}
}

func TestRestaurantsUsecase_GetCommentsRestaurantByRestaurants(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockRestaurantsRepo,
	}

	commentRestaurantData, err := useCase.GetCommentsRestaurantByRestaurants(1)
	require.NoError(t, err)

	mockComment := &models.CommentRestaurantUseCase{}
	mockComment = (*models.CommentRestaurantUseCase)(data.CommentRestaurant)

	if !reflect.DeepEqual(commentRestaurantData.Comment[0], *mockComment) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData.Comment[0], mockComment)
		return
	}
}

func TestRestaurantsUsecase_AddCommentsRestaurantByRestaurants(t *testing.T) {
	mockRestaurantsRepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockRestaurantsRepo,
	}

	item := &models.AddCommentRestaurantUseCase{
		Restaurant: 1,
		User_id: 1,
		Comment_text: "text",
		Comment_rating: 5,
	}
	commentRestaurantData, err := useCase.AddCommentsRestaurantByRestaurants(item)
	require.NoError(t, err)

	mockComment := &models.CommentRestaurantUseCase{}
	mockComment = (*models.CommentRestaurantUseCase)(data.CommentRestaurant)

	if !reflect.DeepEqual(commentRestaurantData, mockComment) {
		t.Errorf("results not match, want %v, have %v", commentRestaurantData, mockComment)
		return
	}
}
