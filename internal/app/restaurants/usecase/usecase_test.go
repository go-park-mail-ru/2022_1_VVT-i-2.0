package ucase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/mock"
	"github.com/stretchr/testify/assert"
)

func TestRestaurantsUcase_GetAllRestaurants(t *testing.T) {
	mockRestaurantsRepo := new(mock.RestaurantsRepo)
	ucase := NewRestaurantsUcase(mockRestaurantsRepo)

	restaurantsResp, err := ucase.GetAllRestaurants()
	assert.NoError(t, err)

	mockRestaurant := &models.RestaurantsUcase{Restaurants: []models.RestaurantUcase{{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}}}

	if !reflect.DeepEqual(restaurantsResp, mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restaurantsResp, mockRestaurant)
		return
	}
}

func TestRestaurantsUcase_GetRestaurantsByCategory(t *testing.T) {
	mockRestaurantsRepo := new(mock.RestaurantsRepo)
	ucase := NewRestaurantsUcase(mockRestaurantsRepo)

	restaurantsResp, err := ucase.GetRestaurantsByCategory(models.GetRestaurantByCategoryUcaseReq{Name: "category"})
	assert.NoError(t, err)

	mockRestaurant := &models.RestaurantsUcase{Restaurants: []models.RestaurantUcase{{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}}}

	if !reflect.DeepEqual(restaurantsResp, mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restaurantsResp, mockRestaurant)
		return
	}
}

func TestRestaurantsUcase_GetRestaurantsByQuery(t *testing.T) {
	mockRestaurantsRepo := new(mock.RestaurantsRepo)
	ucase := NewRestaurantsUcase(mockRestaurantsRepo)

	restaurantsResp, err := ucase.GetRestaurantBySearchQuery(models.GetRestaurantBySearchQueryUcaseReq{Query: "суши"})
	assert.NoError(t, err)

	mockRestaurant := &models.RestaurantsUcase{Restaurants: []models.RestaurantUcase{{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "ImagePath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            9,
		ReviewCount:          2,
		UpMinutsToDelivery:   3,
		DownMinutsToDelivery: 2}}}

	if !reflect.DeepEqual(restaurantsResp, mockRestaurant) {
		t.Errorf("results not match, want %v, have %v", restaurantsResp, mockRestaurant)
		return
	}
}
