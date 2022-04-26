package usecase

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)
import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"

func TestRestaurantsUsecase_GetAllRestaurants(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	restData, err := useCase.GetAllRestaurants()
	require.NoError(t, err)

	if !reflect.DeepEqual(restData.Restaurants[0].Id, data.Rest.Id) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Name, data.Rest.Name) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].City, data.Rest.City) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Address, data.Rest.Address) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Image_path, data.Rest.Image_path) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Slug, data.Rest.Slug) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Min_price, data.Rest.Min_price) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Avg_price, data.Rest.Avg_price) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Rating, data.Rest.Rating) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Restaurants[0].Count_rating, data.Rest.Count_rating) {
		t.Errorf("results not match, want %v, have %v", restData.Restaurants[0], data.Rest)
		return
	}
}

func TestRestaurantsUsecase_GetRestaurantBySluf(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	restData, err := useCase.GetRestaurantBySluf("slug")
	require.NoError(t, err)

	if !reflect.DeepEqual(restData.Id, data.Rest.Id) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Name, data.Rest.Name) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.City, data.Rest.City) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Address, data.Rest.Address) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Image_path, data.Rest.Image_path) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Slug, data.Rest.Slug) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Min_price, data.Rest.Min_price) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Avg_price, data.Rest.Avg_price) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Rating, data.Rest.Rating) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
	if !reflect.DeepEqual(restData.Count_rating, data.Rest.Count_rating) {
		t.Errorf("results not match, want %v, have %v", restData, data.Rest)
		return
	}
}

func TestRestaurantsUsecase_GetDishByRestaurant(t *testing.T) {
	mockAuthorrepo := new(interfaces.RestaurantsRepository)
	useCase := RestaurantsUsecase{
		RestaurantsRepo: mockAuthorrepo,
	}

	dishData, err := useCase.GetDishByRestaurant(1)
	require.NoError(t, err)

	if !reflect.DeepEqual(dishData.Dishes[0].Id, data.Dish.Id) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Restaurant, data.Dish.Restaurant) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Name, data.Dish.Name) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Description, data.Dish.Description) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Image_path, data.Dish.Image_path) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Calories, data.Dish.Calories) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Price, data.Dish.Price) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
	if !reflect.DeepEqual(dishData.Dishes[0].Weight, data.Dish.Weight) {
		t.Errorf("results not match, want %v, have %v", dishData.Dishes[0], data.Rest)
		return
	}
}
