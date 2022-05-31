package repository

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCommentsRepo_GetRestaurantComments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := ReviewsRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"author", "text", "stars", "get_ru_date"})

	expect := &models.GetRestaurantReviewsRepoResp{
		Reviews: make([]models.RestaurantReviewRepo, 3),
	}

	expect.Reviews[0] = models.RestaurantReviewRepo{Author: "author1", Text: "hello", Stars: 1, Date: "today"}
	expect.Reviews[1] = models.RestaurantReviewRepo{Author: "author2", Text: "hello", Stars: 2, Date: "today"}
	expect.Reviews[2] = models.RestaurantReviewRepo{Author: "author3", Text: "hello", Stars: 3, Date: "today"}

	for _, item := range expect.Reviews {
		rows = rows.AddRow(item.Author, item.Text, item.Stars, item.Date)
	}

	// good query
	mock.
		ExpectQuery(`SELECT author, text, stars,...`).
		WithArgs("slug").
		WillReturnRows(rows)
	item, err := repo.GetRestaurantReviews(&models.GetRestaurantReviewsRepoReq{Slug: "slug"})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := &models.GetRestaurantReviewsRepoResp{
		Reviews: make([]models.RestaurantReviewRepo, 3),
	}

	expectResp.Reviews[0] = models.RestaurantReviewRepo{Author: "author1", Text: "hello", Stars: 1, Date: "today"}
	expectResp.Reviews[1] = models.RestaurantReviewRepo{Author: "author2", Text: "hello", Stars: 2, Date: "today"}
	expectResp.Reviews[2] = models.RestaurantReviewRepo{Author: "author3", Text: "hello", Stars: 3, Date: "today"}

	if !reflect.DeepEqual(item.Reviews, expectResp.Reviews) {
		t.Errorf("results not match,\n want %v,\n have %v", item, expectResp.Reviews)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT author, text, stars,...`).
		WithArgs("slug").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantReviews(&models.GetRestaurantReviewsRepoReq{Slug: "slug"})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	mock.
		ExpectQuery(`SELECT author, text, stars,...`).
		WithArgs("slug").
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetRestaurantReviews(&models.GetRestaurantReviewsRepoReq{Slug: "slug"})
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

	repo := ReviewsRepo{
		DB: sqlxDB,
	}

	testReview := &models.AddRestaurantReviewRepoReq{
		Slug:   "slug",
		UserId: 1,
		Text:   "address",
		Rating: 3,
	}

	rows := sqlmock.
		NewRows([]string{"author", "text", "stars", "get_ru_date"})
	expect := &models.RestaurantReviewRepo{
		Author: "author", Text: "text", Stars: 1, Date: "date",
	}
	rows = rows.AddRow(expect.Author, expect.Text, expect.Stars, expect.Date)

	// good query
	mock.
		ExpectQuery(`INSERT`).
		WithArgs(testReview.Slug, testReview.UserId, testReview.Text, testReview.Rating).
		WillReturnRows(rows)
	item, err := repo.AddRestaurantReview(testReview)
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
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", item, expect)
		return
	}

	// query error
	mock.
		ExpectQuery(`INSERT`).
		WithArgs(testReview.Slug, testReview.UserId, testReview.Text, testReview.Rating).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.AddRestaurantReview(testReview)

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
		ExpectQuery(`INSERT`).
		WithArgs(testReview.Slug, testReview.UserId, testReview.Text, testReview.Rating).
		WillReturnError(sql.ErrNoRows)
		//WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.AddRestaurantReview(testReview)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
