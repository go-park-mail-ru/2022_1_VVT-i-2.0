package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewUserRepo(sqlxDB)

	rows := sqlmock.
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataRepo{
		{Id: 1, Name: "Sergey", Phone: "89166152595", Email: "seregey@mail.ru", Avatar: sql.NullString{String: "avatar", Valid: true}},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email, item.Avatar.String)
	}
	fullUpdataData := &models.UpdateUserStorage{Id: int64(1), Name: "name", Email: "email", Avatar: "avatar"}

	// good query
	mock.
		ExpectQuery(`UPDATE`).
		WithArgs(fullUpdataData.Name, fullUpdataData.Email, fullUpdataData.Avatar, fullUpdataData.Id).
		WillReturnRows(rows)
	item, err := repo.UpdateUser(fullUpdataData)
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
		ExpectQuery(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(fullUpdataData.Name, fullUpdataData.Email, fullUpdataData.Avatar, fullUpdataData.Id).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.UpdateUser(fullUpdataData)
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
		ExpectQuery(`UPDATE`).
		WithArgs(fullUpdataData.Name, fullUpdataData.Email, fullUpdataData.Avatar, fullUpdataData.Id).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.UpdateUser(fullUpdataData)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := UserRepo{
		DB: sqlxDB,
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
		ExpectQuery(`SELECT id, phone, email, name, avatar FROM users WHERE id`).
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
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery(`SELECT id, phone, email, name, avatar FROM users WHERE id`).
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
		ExpectQuery(`SELECT id, phone, email, name, avatar FROM users WHERE id`).
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
