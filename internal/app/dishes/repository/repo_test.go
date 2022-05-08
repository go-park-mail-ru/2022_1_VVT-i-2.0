package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestDishesRepo_GetRestaurantBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := DishesRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "image_path", "slug", "min_price", "agg_rating", "review_count", "up_time_to_delivery", "down_time_to_delivery"})
	expect := []*models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.ImagePath, item.Slug, item.MinPrice, item.AggRating, item.ReviewCount, item.UpMinutsToDelivery, item.DownMinutsToDelivery)
	}

	// good query
	mock.
		ExpectQuery("SELECT id, name, image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants WHERE").
		WithArgs("slug").
		WillReturnRows(rows)
	item, err := repo.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq{Slug: "slug"})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, name, image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants WHERE").
		WithArgs("slug").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq{Slug: "slug"})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	mock.
		ExpectQuery("SELECT id, name, image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants WHERE").
		WithArgs("slug").
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq{Slug: "slug"})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestRestaurantsRepo_GetDishByRestaurants(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := DishesRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "restaurant_id", "name", "description", "image_path", "calories", "price", "weight"})
	expect := []*models.DishRepo{
		{Id: 1, RestaurantId: 1, Name: "name", Description: "description", ImagePath: "image_path", Calories: 1, Price: 1, Weight: 1},
		{Id: 2, RestaurantId: 1, Name: "name2", Description: "description2", ImagePath: "image_path2", Calories: 2, Price: 2, Weight: 2},
		{Id: 3, RestaurantId: 1, Name: "name3", Description: "description3", ImagePath: "image_path3", Calories: 3, Price: 3, Weight: 3},
		{Id: 4, RestaurantId: 1, Name: "name4", Description: "description4", ImagePath: "image_path4", Calories: 4, Price: 4, Weight: 4},
		{Id: 5, RestaurantId: 1, Name: "name5", Description: "description5", ImagePath: "image_path5", Calories: 5, Price: 5, Weight: 5},
		{Id: 6, RestaurantId: 1, Name: "name6", Description: "description6", ImagePath: "image_path6", Calories: 6, Price: 6, Weight: 6},
		{Id: 7, RestaurantId: 1, Name: "name7", Description: "description7", ImagePath: "image_path7", Calories: 7, Price: 7, Weight: 7},
		{Id: 8, RestaurantId: 1, Name: "name8", Description: "description8", ImagePath: "image_path8", Calories: 8, Price: 8, Weight: 8},
		{Id: 9, RestaurantId: 1, Name: "name9", Description: "description9", ImagePath: "image_path9", Calories: 9, Price: 9, Weight: 9},
		{Id: 10, RestaurantId: 1, Name: "name10", Description: "description10", ImagePath: "image_path10", Calories: 10, Price: 10, Weight: 10},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.RestaurantId, item.Name, item.Description, item.ImagePath, item.Calories, item.Price, item.Weight)
	}

	// good query
	mock.
		ExpectQuery("SELECT id, restaurant_id, name, description, image_path, calories, price, weight FROM dishes WHERE").
		WithArgs(1).
		WillReturnRows(rows)
	resp, err := repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: 1})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := models.GetRestaurantDishesRepoResp{
		Dishes: []models.DishRepo{{Id: 1, RestaurantId: 1, Name: "name", Description: "description", ImagePath: "image_path", Calories: 1, Price: 1, Weight: 1},
			{Id: 2, RestaurantId: 1, Name: "name2", Description: "description2", ImagePath: "image_path2", Calories: 2, Price: 2, Weight: 2},
			{Id: 3, RestaurantId: 1, Name: "name3", Description: "description3", ImagePath: "image_path3", Calories: 3, Price: 3, Weight: 3},
			{Id: 4, RestaurantId: 1, Name: "name4", Description: "description4", ImagePath: "image_path4", Calories: 4, Price: 4, Weight: 4},
			{Id: 5, RestaurantId: 1, Name: "name5", Description: "description5", ImagePath: "image_path5", Calories: 5, Price: 5, Weight: 5},
			{Id: 6, RestaurantId: 1, Name: "name6", Description: "description6", ImagePath: "image_path6", Calories: 6, Price: 6, Weight: 6},
			{Id: 7, RestaurantId: 1, Name: "name7", Description: "description7", ImagePath: "image_path7", Calories: 7, Price: 7, Weight: 7},
			{Id: 8, RestaurantId: 1, Name: "name8", Description: "description8", ImagePath: "image_path8", Calories: 8, Price: 8, Weight: 8},
			{Id: 9, RestaurantId: 1, Name: "name9", Description: "description9", ImagePath: "image_path9", Calories: 9, Price: 9, Weight: 9},
			{Id: 10, RestaurantId: 1, Name: "name10", Description: "description10", ImagePath: "image_path10", Calories: 10, Price: 10, Weight: 10},
		},
	}

	if !reflect.DeepEqual(resp, &expectResp) {
		t.Errorf("results not match, want %v, have %v", expect, resp)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, restaurant_id, name, description, image_path, calories, price, weight FROM dishes WHERE").
		WithArgs(1).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	mock.
		ExpectQuery("SELECT id, restaurant_id, name, description, image_path, calories, price, weight FROM dishes WHERE").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetRestaurantDishes(models.GetRestaurantDishesRepoReq{Id: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
