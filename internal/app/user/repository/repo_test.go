package repository

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"reflect"
	"testing"
)

func TestGetUserByPhone(t *testing.T) {
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
		NewRows([]string{"id", "name", "phone", "email"})
	expect := []*models.UserDataStorage{
		{1, "Sergey", "89166152595", "seregey@mail.ru"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email)
	}

	mock.
		ExpectQuery("SELECT id, phone, email, name FROM users WHERE phone").
		WithArgs("89166152595").
		WillReturnRows(rows)

	item, err := repo.GetUserByPhone("89166152595")
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

	mock.
		ExpectQuery("SELECT id, phone, email, name FROM users WHERE phone").
		WithArgs("89166152595").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.GetUserByPhone("89166152595")
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
		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnRows(rows)

	_, err = repo.GetUserByPhone("89166152595")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	//if err == nil {
	//	t.Errorf("expected error, got nil")
	//	return
	//}
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := UserRepo{
		DB: sqlxDB,
	}

	name := "Sergey"
	phone := "79166152595"
	email := "sergey@mail.ru"
	testItem := &models.UserAddDataStorage{
		Name: name,
		Phone:phone,
		Email: email,
	}

	//ok query
	mock.
		ExpectExec(`INSERT INTO users (name,phone,email) VALUES ($1,$2,$3) RETURNING id`).
		WithArgs(name, phone, email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.AddUser(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 0 {
		t.Errorf("bad id: want %v, have %v", id, 0)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//query error
	mock.
		ExpectExec(`INSERT INTO items`).
		WithArgs(name, phone, email).
		WillReturnError(fmt.Errorf("bad query"))

	// mock.ExpectClose()

	_, err = repo.AddUser(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// result error
	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(name, phone, email).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("bad_result")))

	_, err = repo.AddUser(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// last id error
	mock.
		ExpectExec(`INSERT INTO items`).
		WithArgs(name, phone, email).
		WillReturnResult(sqlmock.NewResult(0, 0))

	_, err = repo.AddUser(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//func TestUpdateUser(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("cant create mock: %s", err)
//	}
//	defer db.Close()
//
//	sqlxDB := sqlx.NewDb(db, "sqlmock")
//
//	repo := UserRepo{
//		DB: sqlxDB,
//	}
//
//	rows := sqlmock.
//		NewRows([]string{"id", "name", "phone", "email"})
//	expect := []*models.UserDataStorage{
//		{1, "Sergey", "89166152595", "seregey@mail.ru"},
//	}
//	for _, item := range expect {
//		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email)
//	}
//
//	mock.
//		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE id`).
//		WithArgs(1).
//		WillReturnRows(rows)
//
//	item, err := repo.GetUserById(1)
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//		return
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//		return
//	}
//	if !reflect.DeepEqual(item, expect[0]) {
//		t.Errorf("results not match, want %v, have %v", expect[0], item)
//		return
//	}
//
//	mock.
//		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE id`).
//		WithArgs(1).
//		WillReturnError(fmt.Errorf("db_error"))
//	_, err = repo.GetUserById(1)
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//		return
//	}
//	if err == nil {
//		t.Errorf("expected error, got nil")
//		return
//	}
//
//	// row scan error
//	rows = sqlmock.NewRows([]string{"id", "name"}).
//		AddRow(1, "name")
//
//	mock.
//		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE id`).
//		WithArgs(1).
//		WillReturnRows(rows)
//
//	_, err = repo.GetUserById(1)
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//		return
//	}
//	//if err == nil {
//	//	t.Errorf("expected error, got nil")
//	//	return
//	//}
//}

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
		NewRows([]string{"id", "name", "phone", "email"})
	expect := []*models.UserDataStorage{
		{1, "Sergey", "89166152595", "seregey@mail.ru"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email)
	}

	mock.
		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE id`).
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

	mock.
		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE id`).
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
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "name")

	mock.
		ExpectQuery(`SELECT id, phone, email, name FROM users WHERE id`).
		WithArgs(1).
		WillReturnRows(rows)

	_, err = repo.GetUserById(1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	//if err == nil {
	//	t.Errorf("expected error, got nil")
	//	return
	//}
}

func TestHasUserByPhone(t *testing.T) {
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
		NewRows([]string{"id", "name", "phone", "email"})
	expect := []*models.UserDataStorage{
		{1, "Sergey", "89166152595", "seregey@mail.ru"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email)
	}

	mock.
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnRows(rows)

	item, err := repo.HasUserByPhone("89166152595")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, true) {
		t.Errorf("results not match, want %v, have %v", true, item)
		return
	}

	mock.
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.HasUserByPhone("89166152595")
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
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnRows(rows)

	_, err = repo.HasUserByPhone("89166152595")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	//if err == nil {
	//	t.Errorf("expected error, got nil")
	//	return
	//}
}