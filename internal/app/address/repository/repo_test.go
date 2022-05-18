package repository

import (
	"database/sql"
	_ "database/sql"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSuggsRepo_SuggestStreet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := AddrRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"name"})
	expect := []*models.GetStreetRepoResp{
		{Name: "Измайловский проспект"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Name)
	}

	// good query
	qArgs := &models.SuggestStreetRepoReq{Street: "Измайл", CityId: 1}
	mock.
		ExpectQuery(`SELECT name FROM streets ...`).
		WithArgs(qArgs.Street+"%", 5).
		WillReturnRows(rows)
	suggs, err := repo.SuggestStreet(qArgs)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	expectResp := &models.SuggestStreetRepoResp{StreetSuggests: []string{"Измайловский проспект"}}

	if !reflect.DeepEqual(suggs, expectResp) {
		t.Errorf("results not match, want %v, have %v", expect[0], suggs)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT name FROM streets ...`).
		WithArgs(qArgs.Street+"%", 5).
		WillReturnError(sql.ErrConnDone)
	suggs, err = repo.SuggestStreet(&models.SuggestStreetRepoReq{Street: "Измайл", CityId: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
	assert.Nil(t, suggs)

	// row scan error
	// rows = sqlmock.NewRows([]string{"id", "name"}).
	// 	AddRow(1, "name")

	mock.
		ExpectQuery(`SELECT name FROM streets ...`).
		WithArgs(qArgs.Street+"%", 5).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.SuggestStreet(&models.SuggestStreetRepoReq{Street: "Измайл", CityId: 1})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
}

func TestSuggsRepo_SuggestHouse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := AddrRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"name"})
	expect := []*models.GetHouseRepoResp{
		{House: "73A"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.House)
	}

	// good query
	qArgs := &models.SuggestHouseRepoReq{House: "7", StreetId: 1}
	mock.
		ExpectQuery(`SELECT house FROM houses WHERE street_id ...`).
		WithArgs(qArgs.StreetId, qArgs.House+"%", 5).
		WillReturnRows(rows)
	suggs, err := repo.SuggestHouse(qArgs)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	expectResp := &models.SuggestHouseRepoResp{HouseSuggests: []string{"73A"}}

	if !reflect.DeepEqual(suggs, expectResp) {
		t.Errorf("results not match, want %v, have %v", expect[0], suggs)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT house FROM houses WHERE street_id ...`).
		WithArgs(qArgs.StreetId, qArgs.House+"%", 5).
		WillReturnError(sql.ErrConnDone)
	suggs, err = repo.SuggestHouse(qArgs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
	assert.Nil(t, suggs)

	// row scan error
	// rows = sqlmock.NewRows([]string{"id", "name"}).
	// 	AddRow(1, "name")

	mock.
		ExpectQuery(`SELECT house FROM houses WHERE street_id ...`).
		WithArgs(qArgs.StreetId, qArgs.House+"%", 5).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.SuggestHouse(qArgs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
}

func TestSuggsRepo_GetHouse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := AddrRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"house"})
	expect := []*models.GetHouseRepoResp{
		{House: "73A"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.House)
	}

	// good query
	qArgs := &models.GetHouseRepoReq{House: "7", StreetId: 1}
	mock.
		ExpectQuery(`SELECT house FROM houses WHERE street_id...`).
		WithArgs(qArgs.StreetId, qArgs.House).
		WillReturnRows(rows)
	suggs, err := repo.GetHouse(qArgs)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	expectResp := &models.GetHouseRepoResp{House: "73A"}

	if !reflect.DeepEqual(suggs, expectResp) {
		t.Errorf("results not match, want %v, have %v", expect[0], suggs)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT house FROM houses WHERE street_id...`).
		WithArgs(qArgs.StreetId, qArgs.House).
		WillReturnError(sql.ErrConnDone)
	suggs, err = repo.GetHouse(qArgs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
	assert.Nil(t, suggs)

	// row scan error
	// rows = sqlmock.NewRows([]string{"id", "name"}).
	// 	AddRow(1, "name")

	mock.
		ExpectQuery(`SELECT house FROM houses WHERE street_id...`).
		WithArgs(qArgs.StreetId, qArgs.House).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetHouse(qArgs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
}

func TestSuggsRepo_GetStreet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := AddrRepo{
		DB: sqlxDB,
	}

	rows := sqlmock.
		NewRows([]string{"name"})
	expect := []*models.GetStreetRepoResp{
		{Name: "Измайловский проспект"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Name)
	}

	// good query
	qArgs := &models.GetStreetRepoReq{Street: "Измайловский прос", CityId: 0}
	mock.
		ExpectQuery(`SELECT id as streetid, name FROM streets WHERE name ILIKE...`).
		WithArgs(qArgs.Street).
		WillReturnRows(rows)
	suggs, err := repo.GetStreet(qArgs)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	expectResp := &models.GetStreetRepoResp{Name: "Измайловский проспект"}

	if !reflect.DeepEqual(suggs, expectResp) {
		t.Errorf("results not match, want %v, have %v", expect[0], suggs)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT id as streetid, name FROM streets WHERE name ILIKE...`).
		WithArgs(qArgs.Street).
		WillReturnError(sql.ErrConnDone)
	suggs, err = repo.GetStreet(qArgs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
	assert.Nil(t, suggs)

	// row scan error
	// rows = sqlmock.NewRows([]string{"id", "name"}).
	// 	AddRow(1, "name")

	mock.
		ExpectQuery(`SELECT id as streetid, name FROM streets WHERE name ILIKE...`).
		WithArgs(qArgs.Street).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetStreet(qArgs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	assert.Error(t, err)
}

func TestSuggsRepo_GetCity(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewAddrRepo(sqlxDB)

	city, err := repo.GetCity("Москва")

	assert.Equal(t, &models.GetCityRepoResp{CityId: 0, Name: "Москва"}, city)
	assert.NoError(t, err)
}

func TestSuggsRepo_GetCityErr(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewAddrRepo(sqlxDB)

	city, err := repo.GetCity("Нетгорода")

	assert.Nil(t, city)
	assert.Error(t, err)
}
