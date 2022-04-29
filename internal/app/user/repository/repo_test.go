package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"reflect"
	"regexp"
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
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataStorage{
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

	// query error
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT id, phone, email, name, avatar FROM users WHERE phone = $1`)).
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
	mock.
		ExpectQuery(`SELECT id, phone, email, name, avatar FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetUserByPhone("89166152595")
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

	repo := UserRepo{
		DB: sqlxDB,
	}

	name := "Sergey"
	phone := "79166152595"
	email := "sergey@mail.ru"
	testItem := &models.UserAddDataStorage{
		Name:  name,
		Phone:  phone,
		Email:  email,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataStorage{
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

func TestAddUser1(t *testing.T) {
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
		Name:  name,
		Phone:  phone,
		Email:  email,
	}
	var expect models.UserId = 1
	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(expect)

	// good query
	mock.
		ExpectQuery(`INSERT INTO users`).
		WithArgs(name, phone, email).
		WillReturnRows(rows)
	item, err := repo.AddUser1(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if item != expect {
		t.Errorf("bad id: want %v, have %v", item, 0)
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

func TestUpdateUser(t *testing.T) {
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
	expect := []*models.UserDataStorage{
		{Id: 1, Name: "Sergey", Phone: "89166152595", Email: "seregey@mail.ru", Avatar: sql.NullString{String: "avatar", Valid: true}},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.Id, item.Name, item.Phone, item.Email, item.Avatar.String)
	}
	test := &models.UpdateUserDataStorage{Id: models.UserId(1), Name: "name", Email: "email", Avatar: "avatar"}

	// good query
	mock.
		ExpectQuery(regexp.QuoteMeta(`UPDATE users SET name=$1, email=$2, avatar=$3 WHERE id=$4 RETURNING id, name, email, phone, avatar`)).
		WithArgs(test.Name, test.Email, test.Avatar, test.Id).
		WillReturnRows(rows)
	item, err := repo.UpdateUser(test)
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
		ExpectQuery(regexp.QuoteMeta(`UPDATE users SET name=$1, email=$2, avatar=$3 WHERE id=$4 RETURNING id, name, email, phone, avatar`)).
		WithArgs(test.Name, test.Email, test.Avatar, test.Id).
		WillReturnError(fmt.Errorf("db_error"))
	_, err = repo.UpdateUser(test)
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
		ExpectQuery(regexp.QuoteMeta(`UPDATE users SET name=$1, email=$2, avatar=$3 WHERE id=$4 RETURNING id, name, email, phone, avatar`)).
		WithArgs(test.Name, test.Email, test.Avatar, test.Id).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.UpdateUser(test)
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
	expect := []*models.UserDataStorage{
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
		NewRows([]string{"id", "name", "phone", "email", "avatar"})
	expect := []*models.UserDataStorage{
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

	// query error
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
	mock.
		ExpectQuery(`SELECT id FROM users WHERE phone`).
		WithArgs("89166152595").
		WillReturnError(sql.ErrNoRows)
	_, err = repo.HasUserByPhone("89166152595")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}