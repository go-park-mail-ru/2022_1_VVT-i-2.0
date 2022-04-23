package repository

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetRestaurants1(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := RestaurantsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "city", "address", "image_path", "slug", "min_price", "avg_price", "rating", "count_rating"})
	expect := []*models.RestaurantDataStorage{
		{1, "name", "city", "address", "image_path", "slug", 101, 101, 3, 1},
		{2, "name2", "city2", "address2", "image_path2", "slug2", 102, 102, 3, 1},
		{3, "name3", "city3", "address3", "image_path3", "slug3", 103, 103, 3, 1},
		{4, "name4", "city4", "address4", "image_path4", "slug4", 104, 104, 3, 1},
		{5, "name5", "city5", "address5", "image_path5", "slug5", 105, 105, 3, 1},
		{6, "name6", "city6", "address6", "image_path6", "slug6", 106, 106, 3, 1},
		{7, "name7", "city7", "address7", "image_path7", "slug7", 107, 107, 3, 1},
		{8, "name8", "city8", "address8", "image_path8", "slug8", 108, 108, 3, 1},
		{9, "name9", "city9", "address9", "image_path9", "slug9", 109, 109, 3, 1},
		{10, "name10", "city10", "address10", "image_path10", "slug10", 110, 110, 3, 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.City, item.Address, item.Image_path, item.Slug, item.Min_price, item.Avg_price, item.Rating, item.Count_rating)
	}

	t.Run("good query", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants").
			WillReturnRows(rows)

		item, err := repo.GetRestaurants()

		if err != nil {
			t.Errorf("unexpected err: %s", err)
			return
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if !reflect.DeepEqual(item, expect) {
			t.Errorf("results not match, want %v, have %v", expect, item)
			return
		}
	})
	t.Run("query error", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants").
			WillReturnError(fmt.Errorf("db_error"))
		_, err = repo.GetRestaurants()
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
	t.Run("row scan error", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants").
			//WillReturnRows(rows)
			WillReturnError(sql.ErrNoRows)

		_, err = repo.GetRestaurants()
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
}

func TestGetRestaurantsBySlug1(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := RestaurantsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "city", "address", "image_path", "slug", "min_price", "avg_price", "rating", "count_rating"})
	expect := []*models.RestaurantDataStorage{
		{1, "name", "city", "address", "image_path", "slug", 101, 101, 3, 1},
		{2, "name2", "city2", "address2", "image_path2", "slug2", 102, 102, 3, 1},
		{3, "name3", "city3", "address3", "image_path3", "slug3", 103, 103, 3, 1},
		{4, "name4", "city4", "address4", "image_path4", "slug4", 104, 104, 3, 1},
		{5, "name5", "city5", "address5", "image_path5", "slug5", 105, 105, 3, 1},
		{6, "name6", "city6", "address6", "image_path6", "slug6", 106, 106, 3, 1},
		{7, "name7", "city7", "address7", "image_path7", "slug7", 107, 107, 3, 1},
		{8, "name8", "city8", "address8", "image_path8", "slug8", 108, 108, 3, 1},
		{9, "name9", "city9", "address9", "image_path9", "slug9", 109, 109, 3, 1},
		{10, "name10", "city10", "address10", "image_path10", "slug10", 110, 110, 3, 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.City, item.Address, item.Image_path, item.Slug, item.Min_price, item.Avg_price, item.Rating, item.Count_rating)
	}

	t.Run("good query", func(t *testing.T) {
		// good query
		mock.
			ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants WHERE").
			WithArgs("slug").
			WillReturnRows(rows)
		item, err := repo.GetRestaurantsBySlug("slug")
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
	})
	t.Run("query error", func(t *testing.T) {
		// query error
		mock.
			ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants WHERE").
			WithArgs("slug").
			WillReturnError(fmt.Errorf("db_error"))
		_, err = repo.GetRestaurantsBySlug("slug")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
	t.Run("row scan error", func(t *testing.T) {
		// row scan error
		rows = sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "name")

		mock.
			ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants WHERE").
			WithArgs("slug").
			WillReturnError(sql.ErrNoRows)

		_, err = repo.GetRestaurantsBySlug("slug")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
}

func TestGetDishByRestaurants1(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := RestaurantsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "restaurant", "name", "description", "image_path", "calories", "price", "weight"})
	expect := []*models.DishDataStorage{
		{1, 1, "name", "description", "image_path", 1, 1, 1},
		{2, 1, "name2", "description2", "image_path2", 2, 2, 2},
		{3, 1, "name3", "description3", "image_path3", 3, 3, 3},
		{4, 1, "name4", "description4", "image_path4", 4, 4, 4},
		{5, 1, "name5", "description5", "image_path5", 5, 5, 5},
		{6, 1, "name6", "description6", "image_path6", 6, 6, 6},
		{7, 1, "name7", "description7", "image_path7", 7, 7, 7},
		{8, 1, "name8", "description8", "image_path8", 8, 8, 8},
		{9, 1, "name9", "description9", "image_path9", 9,9, 9},
		{10, 1, "name10", "description10", "image_path10", 10, 10, 10},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Restaurant, item.Name, item.Description, item.Image_path, item.Calories, item.Price, item.Weight)
	}

	t.Run("good query", func(t *testing.T) {
		// good query
		mock.
			ExpectQuery("SELECT id, restaurant, name, description, image_path, calories, price, weight FROM dish WHERE").
			WithArgs(1).
			WillReturnRows(rows)
		item, err := repo.GetDishByRestaurants(1)
		if err != nil {
			t.Errorf("unexpected err: %s", err)
			return
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if !reflect.DeepEqual(item, expect) {
			t.Errorf("results not match, want %v, have %v", expect, item)
			return
		}
	})
	t.Run("query error", func(t *testing.T) {
		// query error
		mock.
			ExpectQuery("SELECT id, restaurant, name, description, image_path, calories, price, weight FROM dish WHERE").
			WithArgs(1).
			WillReturnError(fmt.Errorf("db_error"))
		_, err = repo.GetDishByRestaurants(1)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
	t.Run("row scan error", func(t *testing.T) {
		// row scan error
		rows = sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "name")
		mock.
			ExpectQuery("SELECT id, restaurant, name, description, image_path, calories, price, weight FROM dish WHERE").
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)
		_, err = repo.GetDishByRestaurants(1)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
}

func TestGetCommentsRestaurantByRestaurants1(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := RestaurantsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "restaurant", "user_id", "comment_text", "comment_rating"})
	expect := []*models.CommentRestaurantDataStorage{
		{1, 1, 1, "address", 3},
		{2, 1, 2, "address", 4},
		{3, 1, 3, "address", 5},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Restaurant, item.User_id, item.Comment_text, item.Comment_rating)
	}

	t.Run("good query", func(t *testing.T) {
		// good query
		mock.
			ExpectQuery(`SELECT id, restaurant, user_id, comment_text, comment_rating FROM comment_restaurants WHERE id`).
			WithArgs(1).
			WillReturnRows(rows)
		item, err := repo.GetCommentsRestaurantByRestaurants(1)
		if err != nil {
			t.Errorf("unexpected err: %s", err)
			return
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if !reflect.DeepEqual(item, expect) {
			t.Errorf("results not match, want %v, have %v", expect, item)
			return
		}
	})
	t.Run("query error", func(t *testing.T) {
		// query error
		mock.
			ExpectQuery(`SELECT id, restaurant, user_id, comment_text, comment_rating FROM comment_restaurants WHERE id`).
			WithArgs(1).
			WillReturnError(fmt.Errorf("db_error"))
		_, err = repo.GetCommentsRestaurantByRestaurants(1)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
	t.Run("row scan error", func(t *testing.T) {
		// row scan error
		rows = sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "name")
		mock.
			ExpectQuery(`SELECT id, restaurant, user_id, comment_text, comment_rating FROM comment_restaurants WHERE id`).
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)
		_, err = repo.GetCommentsRestaurantByRestaurants(1)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
}

func TestAddCommentsRestaurantByRestaurants1(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := RestaurantsRepo{
		DB: sqlxDB,
	}

	testComment := &models.AddCommentRestaurantDataStorage{
		Restaurant: 1,
		User_id: 1,
		Comment_text: "address",
		Comment_rating: 3,
	}

	rows := sqlmock.
		NewRows([]string{"id", "restaurant", "user_id", "comment_text", "comment_rating"})
	expect := []*models.CommentRestaurantDataStorage{
		{1, testComment.Restaurant, testComment.User_id, testComment.Comment_text, testComment.Comment_rating},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Restaurant, item.User_id, item.Comment_text, item.Comment_rating)
	}

	t.Run("good query", func(t *testing.T) {
		// good query
		mock.
			ExpectQuery(`INSERT INTO comment_restaurants`).
			WithArgs(testComment.Restaurant, testComment.User_id, testComment.Comment_text, testComment.Comment_rating).
			WillReturnRows(rows)
		item, err := repo.AddCommentsRestaurantByRestaurants(testComment)
		if err != nil {
			t.Errorf("unexpected err: %s", err)
			return
		}
		if item == nil {
			t.Errorf("bad item")
			return
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if !reflect.DeepEqual(item, expect[0]) {
			t.Errorf("results not match, want %v, have %v", item, expect[0])
			return
		}
	})
	t.Run("query error", func(t *testing.T) {
		// query error
		mock.
			ExpectQuery(`INSERT INTO comment_restaurants`).
			WithArgs(testComment.Restaurant, testComment.User_id, testComment.Comment_text, testComment.Comment_rating).
			WillReturnError(fmt.Errorf("db_error"))
		_, err = repo.AddCommentsRestaurantByRestaurants(testComment)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
	t.Run("row scan error", func(t *testing.T) {
		// row scan error
		mock.
			ExpectQuery(`INSERT INTO comment_restaurants`).
			WithArgs(testComment.Restaurant, testComment.User_id, testComment.Comment_text, testComment.Comment_rating).
			WillReturnError(fmt.Errorf("db_error"))
		_, err = repo.AddCommentsRestaurantByRestaurants(testComment)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
			return
		}
		if err == nil {
			t.Errorf("expected error, got nil")
			return
		}
	})
}

//func Test1(t *testing.T) {
//	t.Parallel()
//
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("cant create mock: %s", err)
//	}
//	defer db.Close()
//
//	sqlxDB := sqlx.NewDb(db, "sqlmock")
//
//	repo := RestaurantsRepo{
//		DB: sqlxDB,
//	}
//
//
//
//	t.Run("good query", func(t *testing.T) {
//
//	})
//	t.Run("query error", func(t *testing.T) {
//
//	})
//	t.Run("row scan error", func(t *testing.T) {
//
//	})
//}