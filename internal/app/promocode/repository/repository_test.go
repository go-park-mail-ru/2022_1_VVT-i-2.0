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

func TestRestaurantsRepo_GetRestaurants(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewPromocodeRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"name", "slug", "text", "discount", "price_reduction", "min_price", "promocode", "image_path", "logo_path"})
	expect := []*models.PromocodeRepoResp{
		{Promocode: "promo1", RestaurantSlug: "slug1", RestaurantName: "name1", MinPrice: 101, Discount: 1, PriceReduction: 2, Text: "text1", LogoImage: "logo1", Image: "image1"},
		{Promocode: "promo2", RestaurantSlug: "slug2", RestaurantName: "name2", MinPrice: 102, Discount: 1, PriceReduction: 2, Text: "text2", LogoImage: "logo2", Image: "image2"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.RestaurantName, item.RestaurantSlug, item.Text, item.Discount, item.PriceReduction, item.MinPrice, item.Promocode, item.Image, item.LogoImage)
	}

	// good query
	mock.
		ExpectQuery(`SELECT r.name, r.slug, p.text, p.discount, p.price_reduction, p.min_price, p.promocode, p.image_path, p.logo_path FROM promocodes p JOIN restaurants r ON p.restaurant_id=r.id`).
		WillReturnRows(rows)
	restaurantsResp, err := repo.GetAllPromocodes()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	expectResp := &models.GetPromocodesRepoResp{Promos: []models.PromocodeRepoResp{
		{Promocode: "promo1", RestaurantSlug: "slug1", RestaurantName: "name1", MinPrice: 101, Discount: 1, PriceReduction: 2, Text: "text1", LogoImage: "logo1", Image: "image1"},
		{Promocode: "promo2", RestaurantSlug: "slug2", RestaurantName: "name2", MinPrice: 102, Discount: 1, PriceReduction: 2, Text: "text2", LogoImage: "logo2", Image: "image2"},
	},
	}

	if !reflect.DeepEqual(restaurantsResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", expectResp, restaurantsResp)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT r.name, r.slug, p.text, p.discount, p.price_reduction, p.min_price, p.promocode, p.image_path, p.logo_path FROM promocodes p JOIN restaurants r ON p.restaurant_id=r.id`).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetAllPromocodes()
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
		ExpectQuery(`SELECT r.name, r.slug, p.text, p.discount, p.price_reduction, p.min_price, p.promocode, p.image_path, p.logo_path FROM promocodes p JOIN restaurants r ON p.restaurant_id=r.id`).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetAllPromocodes()
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
