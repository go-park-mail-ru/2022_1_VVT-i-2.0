package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetUserByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewAuthRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataRepo{
		{Id: 1, Name: "Sergey", Phone: "89166152595", Email: "seregey@mail.ru", Avatar: sql.NullString{String: "avatar", Valid: true}},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email, item.Avatar.String)
	}

	// good query
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT id, phone, email, name, avatar FROM users WHERE phone = $1`)).
		WithArgs("89166152595").
		WillReturnRows(rows)
	item, err := repo.GetUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
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
		ExpectQuery(regexp.QuoteMeta(`SELECT id, phone, email, name, avatar FROM users WHERE phone = $1`)).
		WithArgs("89166152595").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
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
		ExpectQuery(`SELECT id, phone, email, name, avatar FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewAuthRepo(sqlxDB)

	name := "Sergey"
	phone := "79166152595"
	email := "sergey@mail.ru"
	testItem := &models.AddUserRepoReq{
		Name:  name,
		Phone: phone,
		Email: email,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataRepo{
		{Id: 1, Name: "Sergey", Phone: "89166152595", Email: "seregey@mail.ru", Avatar: sql.NullString{String: "avatar", Valid: true}},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email, item.Avatar.String)
	}

	// good query
	mock.
		ExpectQuery(`INSERT INTO users`).
		WithArgs(name, phone, email).
		WillReturnRows(rows)
	item, err := repo.AddUser(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if item == nil {
		t.Errorf("bad id: want %v, have %v", item, 0)
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
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO users`)).
		WithArgs(name, phone, email).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.AddUser(testItem)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	// row scan error
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO users`)).
		WithArgs(name, phone, email).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.AddUser(testItem)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestHasUserByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewAuthRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataRepo{
		{Id: 1, Name: "Sergey", Phone: "89166152595", Email: "seregey@mail.ru", Avatar: sql.NullString{String: "avatar", Valid: true}},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email, item.Avatar.String)
	}

	// good query
	mock.
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnRows(rows)
	item, err := repo.HasUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, models.HasSuchUserRepoResp{IsRegistered: true}) {
		t.Errorf("results not match, want %v, have %v", true, item)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.HasUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
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
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnError(sql.ErrNoRows)
	_, err = repo.HasUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetTopUserAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewAuthRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"address"})
	expect := []*models.GetTopUserAddrRepoResp{
		{Address: "Москва, Петровка, 38"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Address)
	}

	// good query
	mock.
		ExpectQuery(`SELECT address FROM `).
		WithArgs("1").
		WillReturnRows(rows)

	item, err := repo.GetTopUserAddr(&models.GetTopUserAddrRepoReq{UserId: 1})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(item, &models.GetTopUserAddrRepoResp{}) {
		t.Errorf("results not match, want %v, have %v", true, item)
		return
	}

	// // query error
	// mock.
	// 	ExpectQuery(`SELECT id FROM users WHERE phone`).
	// 	WithArgs("89166152595").
	// 	WillReturnError(fmt.Errorf("db_error"))
	// _, err = repo.HasUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// 	return
	// }
	// if err == nil {
	// 	t.Errorf("expected error, got nil")
	// 	return
	// }

	// // row scan error
	// mock.
	// 	ExpectQuery(`SELECT id FROM users WHERE phone`).
	// 	WithArgs("89166152595").
	// 	WillReturnError(sql.ErrNoRows)
	// _, err = repo.HasUserByPhone(models.UserByPhoneRepoReq{Phone: "89166152595"})
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// 	return
	// }
	// if err != nil {
	// 	t.Errorf("expected error, got nil")
	// 	return
	// }
}
