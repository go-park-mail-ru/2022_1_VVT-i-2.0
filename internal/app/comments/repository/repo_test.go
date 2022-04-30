package repository

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCommentsRepo_GetRestaurantByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := CommentsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "city", "address", "image_path", "slug", "min_price", "avg_price", "rating", "count_rating"})
	expect := []*models.RestaurantDataStorage{
		{1, "name", "city", "address", "image_path", "slug", 101, 101, 3, 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.City, item.Address, item.ImagePath, item.Slug, item.MinPrice, item.Avg_price, item.Rating, item.CountRating)
	}

	// good query
	mock.
		ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants WHERE").
		WithArgs(1).
		WillReturnRows(rows)
	item, err := repo.GetRestaurantByID(1)
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
		ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants WHERE").
		WithArgs("slug").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantBySlug("slug")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "name")
	mock.
		ExpectQuery("SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating FROM restaurants WHERE").
		WithArgs("slug").
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetRestaurantBySlug("slug")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCommentsRepo_GetRestaurantComments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := CommentsRepo{
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

	// good query
	mock.
		ExpectQuery("SELECT").
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

	// query error
	mock.
		ExpectQuery(`SELECT`).
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

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "name")
	mock.
		ExpectQuery(`SELECT`).
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
}

func TestCommentsRepo_AddRestaurantComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := CommentsRepo{
		DB: sqlxDB,
	}

	testComment := &models.AddCommentRestaurantDataStorage{
		Restaurant:     1,
		User_id:        1,
		Comment_text:   "address",
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
}

func TestCommentRepo_UpdateRestaurantRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := CommentsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "city", "address", "image_path", "slug", "min_price", "avg_price", "rating", "count_rating"})
	expect := []*models.RestaurantDataStorage{
		{1, "name", "city", "address", "image_path", "slug", 101, 101, 10, 2},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.City, item.Address, item.ImagePath, item.Slug, item.MinPrice, item.Avg_price, item.Rating, item.CountRating)
	}

	// good query
	mock.
		ExpectQuery(`UPDATE`).
		WithArgs(10, 2, 1).
		WillReturnRows(rows)
	item, err := repo.UpdateRestaurantRating(1, 10, 2)
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
		ExpectQuery(`UPDATE`).
		WithArgs(10, 2, 1).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.UpdateRestaurantRating(1, 10, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "name")
	mock.
		ExpectQuery(`UPDATE`).
		WithArgs(10, 2, 1).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.UpdateRestaurantRating(1, 10, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
