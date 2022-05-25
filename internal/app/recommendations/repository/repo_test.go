package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRecommendationsRepo_GetRestaurantDishesRepo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := RecommendationsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "restaurant_id", "category", "name", "description", "image_path", "calories", "price", "weight"})
	expect := []*models.DishCategoriesRepo{
		{Id: 1, RestaurantId: 1, Category: 1, Name: "name", Description: "description", ImagePath: "image_path", Calories: 1, Price: 1, Weight: 1},
		{Id: 2, RestaurantId: 1, Category: 1, Name: "name2", Description: "description2", ImagePath: "image_path2", Calories: 2, Price: 2, Weight: 2},
		{Id: 3, RestaurantId: 1, Category: 1, Name: "name3", Description: "description3", ImagePath: "image_path3", Calories: 3, Price: 3, Weight: 3},
		{Id: 4, RestaurantId: 1, Category: 1, Name: "name4", Description: "description4", ImagePath: "image_path4", Calories: 4, Price: 4, Weight: 4},
		{Id: 5, RestaurantId: 1, Category: 1, Name: "name5", Description: "description5", ImagePath: "image_path5", Calories: 5, Price: 5, Weight: 5},
		{Id: 6, RestaurantId: 1, Category: 1, Name: "name6", Description: "description6", ImagePath: "image_path6", Calories: 6, Price: 6, Weight: 6},
		{Id: 7, RestaurantId: 1, Category: 1, Name: "name7", Description: "description7", ImagePath: "image_path7", Calories: 7, Price: 7, Weight: 7},
		{Id: 8, RestaurantId: 1, Category: 1, Name: "name8", Description: "description8", ImagePath: "image_path8", Calories: 8, Price: 8, Weight: 8},
		{Id: 9, RestaurantId: 1, Category: 1, Name: "name9", Description: "description9", ImagePath: "image_path9", Calories: 9, Price: 9, Weight: 9},
		{Id: 10, RestaurantId: 1, Category: 1, Name: "name10", Description: "description10", ImagePath: "image_path10", Calories: 10, Price: 10, Weight: 10},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Category, item.RestaurantId, item.Name, item.Description, item.ImagePath, item.Calories, item.Price, item.Weight)
	}

	// good query
	mock.
		ExpectQuery("SELECT").
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

	expectResp := []models.DishCategoriesRepo{
		{Id: 1, RestaurantId: 1, Category: 1, Name: "name", Description: "description", ImagePath: "image_path", Calories: 1, Price: 1, Weight: 1},
		{Id: 2, RestaurantId: 1, Category: 1, Name: "name2", Description: "description2", ImagePath: "image_path2", Calories: 2, Price: 2, Weight: 2},
		{Id: 3, RestaurantId: 1, Category: 1, Name: "name3", Description: "description3", ImagePath: "image_path3", Calories: 3, Price: 3, Weight: 3},
		{Id: 4, RestaurantId: 1, Category: 1, Name: "name4", Description: "description4", ImagePath: "image_path4", Calories: 4, Price: 4, Weight: 4},
		{Id: 5, RestaurantId: 1, Category: 1, Name: "name5", Description: "description5", ImagePath: "image_path5", Calories: 5, Price: 5, Weight: 5},
		{Id: 6, RestaurantId: 1, Category: 1, Name: "name6", Description: "description6", ImagePath: "image_path6", Calories: 6, Price: 6, Weight: 6},
		{Id: 7, RestaurantId: 1, Category: 1, Name: "name7", Description: "description7", ImagePath: "image_path7", Calories: 7, Price: 7, Weight: 7},
		{Id: 8, RestaurantId: 1, Category: 1, Name: "name8", Description: "description8", ImagePath: "image_path8", Calories: 8, Price: 8, Weight: 8},
		{Id: 9, RestaurantId: 1, Category: 1, Name: "name9", Description: "description9", ImagePath: "image_path9", Calories: 9, Price: 9, Weight: 9},
		{Id: 10, RestaurantId: 1, Category: 1, Name: "name10", Description: "description10", ImagePath: "image_path10", Calories: 10, Price: 10, Weight: 10},
	}

	if !reflect.DeepEqual(resp.Dishes, expectResp) {
		t.Errorf("results not match,\n want %v,\n have %v", resp.Dishes, expectResp)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
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
		ExpectQuery("SELECT").
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
