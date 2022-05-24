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

	repo := NewCommentsRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "slug", "image_path", "min_price", "up_time_to_delivery", "down_time_to_delivery", "review_count", "agg_rating"})
	expect := &models.RestaurantRepo{Id: 1, Name: "name", ImagePath: "image_path", Slug: "slug1", MinPrice: 100, AggRating: 100, ReviewCount: 100, UpMinutsToDelivery: 100, DownMinutsToDelivery: 100,}
	rows = rows.AddRow(expect.Id, expect.Name, expect.Slug, expect.ImagePath, expect.MinPrice, expect.AggRating, expect.ReviewCount, expect.UpMinutsToDelivery, expect.DownMinutsToDelivery)

	// good query
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnRows(rows)
	item, err := repo.GetRestaurantByID(models.GetRestaurantByIdRepoReq{Id: 1})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", item, expect)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantByID(models.GetRestaurantByIdRepoReq{Id: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	//rows = sqlmock.NewRows([]string{"id", "name"}).
	//	AddRow(1, "name")
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetRestaurantByID(models.GetRestaurantByIdRepoReq{Id: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCommentsRepo_GetRestaurantBySlug(t *testing.T) {
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
		NewRows([]string{"id", "name", "slug", "image_path", "min_price", "up_time_to_delivery", "down_time_to_delivery", "review_count", "agg_rating"})
	expect := &models.RestaurantRepo{Id: 1, Name: "name", ImagePath: "image_path", Slug: "slug1", MinPrice: 100, AggRating: 100, ReviewCount: 100, UpMinutsToDelivery: 100, DownMinutsToDelivery: 100}
	rows = rows.AddRow(expect.Id, expect.Name, expect.Slug, expect.ImagePath, expect.MinPrice, expect.AggRating, expect.ReviewCount, expect.UpMinutsToDelivery, expect.DownMinutsToDelivery)

	// good query
	mock.
		ExpectQuery(`SELECT`).
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
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", item, expect)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT`).
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
	//rows = sqlmock.NewRows([]string{"id", "name"}).
	//	AddRow(1, "name")
	mock.
		ExpectQuery(`SELECT`).
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
		NewRows([]string{"restaurant_id", "author", "text", "stars", "get_ru_date"})
	//expect := []*models.CommentRestaurantDataStorage{
	//	{1, "author1", "hello", 1, "today"},
	//	{2, "author2", "hello", 2, "today"},
	//	{3, "author3", "hello", 3, "today"},
	//}

	//expect := &models.CommentsRestaurantDataStorage{
	//	Comments: []models.CommentRestaurantDataStorage{
	//		{1, "author1", "hello", 1, "today"},
	//		{2, "author2", "hello", 2, "today"},
	//		{3, "author3", "hello", 3, "today"},
	//
	//}}

	expect := &models.CommentsRestaurantDataStorage{
		Comments: make([]models.CommentRestaurantDataStorage, 3),
	}

	expect.Comments[0] = models.CommentRestaurantDataStorage{RestaurantId: 1, Author: "author1", Text: "hello", Stars: 1, Date: "today"}
	expect.Comments[1] = models.CommentRestaurantDataStorage{RestaurantId: 2, Author: "author2", Text: "hello", Stars: 2, Date: "today"}
	expect.Comments[2] = models.CommentRestaurantDataStorage{RestaurantId: 3, Author: "author3", Text: "hello", Stars: 3, Date: "today"}

	for _, item := range expect.Comments {
		rows = rows.AddRow(item.RestaurantId, item.Author, item.Text, item.Stars, item.Date)
	}

	// good query
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnRows(rows)
	item, err := repo.GetRestaurantComments(models.GetRestaurantCommentsRepoReq{Id: 1})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := &models.CommentsRestaurantDataStorage{
		Comments: make([]models.CommentRestaurantDataStorage, 3),
	}

	expectResp.Comments[0] = models.CommentRestaurantDataStorage{RestaurantId: 1, Author: "author1", Text: "hello", Stars: 1, Date: "today"}
	expectResp.Comments[1] = models.CommentRestaurantDataStorage{RestaurantId: 2, Author: "author2", Text: "hello", Stars: 2, Date: "today"}
	expectResp.Comments[2] = models.CommentRestaurantDataStorage{RestaurantId: 3, Author: "author3", Text: "hello", Stars: 3, Date: "today"}

	if !reflect.DeepEqual(item.Comments, expectResp.Comments) {
		t.Errorf("results not match,\n want %v,\n have %v", item, expectResp.Comments)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantComments(models.GetRestaurantCommentsRepoReq{Id: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	//rows = sqlmock.NewRows([]string{"id", "name"}).
	//	AddRow(1, "name")
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetRestaurantComments(models.GetRestaurantCommentsRepoReq{Id: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCommentsRepo_GetUserById(t *testing.T) {
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
		NewRows([]string{"id", "phone", "email", "name", "avatar"})
	expect := []*models.UserDataRepo{
		{Id: 1, Name: "Sergey", Phone: "89166152595", Email: "seregey@mail.ru", Avatar: sql.NullString{String: "avatar", Valid: true}},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email, item.Avatar.String)
	}

	// good query
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnRows(rows)
	item, err := repo.GetUserById(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item.Id, expect[0].Id) {
		t.Errorf("results not match, want %v, have %v", item, expect[0])
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetUserById(1)
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
		ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetUserById(1)
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

	testComment := &models.AddRestaurantCommentRepoReq{
		RestaurantId:  1,
		User:          "author",
		CommentText:   "address",
		CommentRating: 3,
	}

	rows := sqlmock.
		NewRows([]string{"restaurant_id", "author", "text", "stars", "get_ru_date"})
	expect := &models.CommentRestaurantDataStorage{
		RestaurantId: testComment.RestaurantId, Author: testComment.User, Text: testComment.CommentText, Stars: testComment.CommentRating, Date: "date",
	}
	rows = rows.AddRow(expect.RestaurantId, expect.Author, expect.Text, expect.Stars, expect.Date)

	// good query
	mock.
		ExpectQuery(`INSERT`).
		WithArgs(testComment.RestaurantId, testComment.User, testComment.CommentText, testComment.CommentRating).
		WillReturnRows(rows)
	item, err := repo.AddRestaurantComment(models.AddRestaurantCommentRepoReq{
		RestaurantId: testComment.RestaurantId,
		User: testComment.User,
		CommentText: testComment.CommentText,
		CommentRating: testComment.CommentRating,
	})
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
		WithArgs(testComment.RestaurantId, testComment.User, testComment.CommentText, testComment.CommentRating).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.AddRestaurantComment(models.AddRestaurantCommentRepoReq{
		RestaurantId: testComment.RestaurantId,
		User: testComment.User,
		CommentText: testComment.CommentText,
		CommentRating: testComment.CommentRating,
	})
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
		WithArgs(testComment.RestaurantId, testComment.User, testComment.CommentText, testComment.CommentRating).
		WillReturnError(sql.ErrNoRows)
		//WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.AddRestaurantComment(models.AddRestaurantCommentRepoReq{
		RestaurantId: testComment.RestaurantId,
		User: testComment.User,
		CommentText: testComment.CommentText,
		CommentRating: testComment.CommentRating,
	})
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
		NewRows([]string{"id", "name", "slug", "image_path", "min_price", "up_time_to_delivery", "down_time_to_delivery", "review_count", "agg_rating"})
	expect := &models.RestaurantRepo{
		Id: 1, Name: "name", ImagePath: "image_path", Slug: "slug1", MinPrice: 100, AggRating: 100, ReviewCount: 100, UpMinutsToDelivery: 100, DownMinutsToDelivery: 100,
	}
	rows = rows.AddRow(expect.Id, expect.Name, expect.Slug, expect.ImagePath, expect.MinPrice, expect.AggRating, expect.ReviewCount, expect.UpMinutsToDelivery, expect.DownMinutsToDelivery)

	// good query
	mock.
		ExpectQuery(`UPDATE`).
		WithArgs(10, 2, 1).
		WillReturnRows(rows)
	item, err := repo.UpdateRestaurantRating(models.UpdateRestaurantRatingRepoReq{
		RestId: 1,
		NewRestRating: 10,
		CountRating: 2,
	})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", item, expect)
		return
	}

	// query error
	mock.
		ExpectQuery(`UPDATE`).
		WithArgs(10, 2, 1).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.UpdateRestaurantRating(models.UpdateRestaurantRatingRepoReq{
		RestId: 1,
		NewRestRating: 10,
		CountRating: 2,
	})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	//rows = sqlmock.NewRows([]string{"id", "name"}).
	//	AddRow(1, "name")
	mock.
		ExpectQuery(`UPDATE`).
		WithArgs(10, 2, 1).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.UpdateRestaurantRating(models.UpdateRestaurantRatingRepoReq{
		RestId: 1,
		NewRestRating: 10,
		CountRating: 2,
	})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
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
