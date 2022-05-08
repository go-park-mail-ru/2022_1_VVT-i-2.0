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

func TestRestaurantsRepo_GetRestaurants(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewRestaurantsRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "image_path", "slug", "min_price", "agg_rating", "review_count", "up_time_to_delivery", "down_time_to_delivery"})
	expect := []*models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 2, Name: "name2", ImagePath: "image_path2", Slug: "slug2", MinPrice: 102, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 3, Name: "name3", ImagePath: "image_path3", Slug: "slug3", MinPrice: 103, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 4, Name: "name4", ImagePath: "image_path4", Slug: "slug4", MinPrice: 104, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 5, Name: "name5", ImagePath: "image_path5", Slug: "slug5", MinPrice: 105, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 6, Name: "name6", ImagePath: "image_path6", Slug: "slug6", MinPrice: 106, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 7, Name: "name7", ImagePath: "image_path7", Slug: "slug7", MinPrice: 107, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 8, Name: "name8", ImagePath: "image_path8", Slug: "slug8", MinPrice: 108, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 9, Name: "name9", ImagePath: "image_path9", Slug: "slug9", MinPrice: 109, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 10, Name: "name10", ImagePath: "image_path10", Slug: "slug10", MinPrice: 110, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.ImagePath, item.Slug, item.MinPrice, item.AggRating, item.ReviewCount, item.UpMinutsToDelivery, item.DownMinutsToDelivery)
	}

	// good query
	mock.
		ExpectQuery("SELECT id, name, image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants").
		WillReturnRows(rows)
	restaurantsResp, err := repo.GetRestaurants()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := &models.RestaurantsRepo{Restaurants: []models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 2, Name: "name2", ImagePath: "image_path2", Slug: "slug2", MinPrice: 102, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 3, Name: "name3", ImagePath: "image_path3", Slug: "slug3", MinPrice: 103, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 4, Name: "name4", ImagePath: "image_path4", Slug: "slug4", MinPrice: 104, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 5, Name: "name5", ImagePath: "image_path5", Slug: "slug5", MinPrice: 105, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 6, Name: "name6", ImagePath: "image_path6", Slug: "slug6", MinPrice: 106, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 7, Name: "name7", ImagePath: "image_path7", Slug: "slug7", MinPrice: 107, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 8, Name: "name8", ImagePath: "image_path8", Slug: "slug8", MinPrice: 108, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 9, Name: "name9", ImagePath: "image_path9", Slug: "slug9", MinPrice: 109, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 10, Name: "name10", ImagePath: "image_path10", Slug: "slug10", MinPrice: 110, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}}

	if !reflect.DeepEqual(restaurantsResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", expectResp, restaurantsResp)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, name, image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants").
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

	// row scan error
	mock.
		ExpectQuery("SELECT id, name, image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants").
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
}

func TestRestaurantsRepo_GetRestaurantsByQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewRestaurantsRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "image_path", "slug", "min_price", "agg_rating", "review_count", "up_time_to_delivery", "down_time_to_delivery"})
	expect := []*models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 2, Name: "name2", ImagePath: "image_path2", Slug: "slug2", MinPrice: 102, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 3, Name: "name3", ImagePath: "image_path3", Slug: "slug3", MinPrice: 103, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 4, Name: "name4", ImagePath: "image_path4", Slug: "slug4", MinPrice: 104, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 5, Name: "name5", ImagePath: "image_path5", Slug: "slug5", MinPrice: 105, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 6, Name: "name6", ImagePath: "image_path6", Slug: "slug6", MinPrice: 106, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 7, Name: "name7", ImagePath: "image_path7", Slug: "slug7", MinPrice: 107, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 8, Name: "name8", ImagePath: "image_path8", Slug: "slug8", MinPrice: 108, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 9, Name: "name9", ImagePath: "image_path9", Slug: "slug9", MinPrice: 109, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 10, Name: "name10", ImagePath: "image_path10", Slug: "slug10", MinPrice: 110, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.ImagePath, item.Slug, item.MinPrice, item.AggRating, item.ReviewCount, item.UpMinutsToDelivery, item.DownMinutsToDelivery)
	}

	// good query
	mock.
		ExpectQuery(`SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE c.name ILIKE `).
		WithArgs("category%").
		WillReturnRows(rows)
	restaurantsResp, err := repo.GetRestaurantsBySeachQuery(models.GetRestaurantBySearchQueryRepoReq{Query: "category"})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := &models.RestaurantsRepo{Restaurants: []models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 2, Name: "name2", ImagePath: "image_path2", Slug: "slug2", MinPrice: 102, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 3, Name: "name3", ImagePath: "image_path3", Slug: "slug3", MinPrice: 103, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 4, Name: "name4", ImagePath: "image_path4", Slug: "slug4", MinPrice: 104, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 5, Name: "name5", ImagePath: "image_path5", Slug: "slug5", MinPrice: 105, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 6, Name: "name6", ImagePath: "image_path6", Slug: "slug6", MinPrice: 106, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 7, Name: "name7", ImagePath: "image_path7", Slug: "slug7", MinPrice: 107, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 8, Name: "name8", ImagePath: "image_path8", Slug: "slug8", MinPrice: 108, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 9, Name: "name9", ImagePath: "image_path9", Slug: "slug9", MinPrice: 109, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 10, Name: "name10", ImagePath: "image_path10", Slug: "slug10", MinPrice: 110, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}}

	if !reflect.DeepEqual(restaurantsResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", expectResp, restaurantsResp)
		return
	}
}

func TestRestaurantsRepo_GetRestaurantsByCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewRestaurantsRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "image_path", "slug", "min_price", "agg_rating", "review_count", "up_time_to_delivery", "down_time_to_delivery"})
	expect := []*models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 2, Name: "name2", ImagePath: "image_path2", Slug: "slug2", MinPrice: 102, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 3, Name: "name3", ImagePath: "image_path3", Slug: "slug3", MinPrice: 103, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 4, Name: "name4", ImagePath: "image_path4", Slug: "slug4", MinPrice: 104, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 5, Name: "name5", ImagePath: "image_path5", Slug: "slug5", MinPrice: 105, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 6, Name: "name6", ImagePath: "image_path6", Slug: "slug6", MinPrice: 106, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 7, Name: "name7", ImagePath: "image_path7", Slug: "slug7", MinPrice: 107, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 8, Name: "name8", ImagePath: "image_path8", Slug: "slug8", MinPrice: 108, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 9, Name: "name9", ImagePath: "image_path9", Slug: "slug9", MinPrice: 109, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 10, Name: "name10", ImagePath: "image_path10", Slug: "slug10", MinPrice: 110, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.ImagePath, item.Slug, item.MinPrice, item.AggRating, item.ReviewCount, item.UpMinutsToDelivery, item.DownMinutsToDelivery)
	}

	// good query
	mock.
		ExpectQuery(`SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE c.name=`).
		WithArgs("category").
		WillReturnRows(rows)
	restaurantsResp, err := repo.GetRestaurantsByCategory(models.GetRestaurantByCategoryRepoReq{Name: "category"})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := &models.RestaurantsRepo{Restaurants: []models.RestaurantRepo{
		{Id: 1, Name: "name1", ImagePath: "image_path1", Slug: "slug1", MinPrice: 101, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 2, Name: "name2", ImagePath: "image_path2", Slug: "slug2", MinPrice: 102, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 3, Name: "name3", ImagePath: "image_path3", Slug: "slug3", MinPrice: 103, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 4, Name: "name4", ImagePath: "image_path4", Slug: "slug4", MinPrice: 104, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 5, Name: "name5", ImagePath: "image_path5", Slug: "slug5", MinPrice: 105, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 6, Name: "name6", ImagePath: "image_path6", Slug: "slug6", MinPrice: 106, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 7, Name: "name7", ImagePath: "image_path7", Slug: "slug7", MinPrice: 107, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 8, Name: "name8", ImagePath: "image_path8", Slug: "slug8", MinPrice: 108, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 9, Name: "name9", ImagePath: "image_path9", Slug: "slug9", MinPrice: 109, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
		{Id: 10, Name: "name10", ImagePath: "image_path10", Slug: "slug10", MinPrice: 110, AggRating: 100, ReviewCount: 25, UpMinutsToDelivery: 1, DownMinutsToDelivery: 1},
	}}

	if !reflect.DeepEqual(restaurantsResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", expectResp, restaurantsResp)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE `).
		WithArgs("category").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetRestaurantsByCategory(models.GetRestaurantByCategoryRepoReq{Name: "category"})
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
		ExpectQuery(`SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE `).
		WithArgs("category").
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetRestaurantsByCategory(models.GetRestaurantByCategoryRepoReq{Name: "category"})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
