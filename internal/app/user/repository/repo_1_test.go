package repository

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/jmoiron/sqlx"
	"reflect"
	"regexp"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetUserByPhone1(t *testing.T) {
	t.Parallel()

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

	t.Run("good query", func(t *testing.T) {
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
	})
	t.Run("query error", func(t *testing.T) {
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
	})
	t.Run("row scan error", func(t *testing.T) {
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
	})
}

func TestAddUser11(t *testing.T) {
	t.Parallel()

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

	t.Run("good query", func(t *testing.T) {
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
	})

	t.Run("query error", func(t *testing.T) {
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
	})

	t.Run("row scan error", func(t *testing.T) {
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
	})
}

func TestAddUser111(t *testing.T) {
	t.Parallel()

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

	t.Run("good query", func(t *testing.T) {
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
	})

	t.Run("query error", func(t *testing.T) {
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
	})

	t.Run("row scan error", func(t *testing.T) {
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
	})
}

func TestUpdateUser1(t *testing.T) {
	t.Parallel()

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
	test := &models.UpdateUserStorage{Id: models.UserId(1), Name: "name", Email: "email", Avatar: "avatar"}

	t.Run("good query", func(t *testing.T) {
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
	})

	t.Run("query error", func(t *testing.T) {
		// query error
		mock.
			ExpectQuery(`UPDATE`).
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
	})

	t.Run("row scan error", func(t *testing.T) {
		// row scan error
		mock.
			ExpectQuery(`UPDATE`).
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
	})
}

func TestGetUserById1(t *testing.T) {
	t.Parallel()

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

	t.Run("good query", func(t *testing.T) {
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
	})

	t.Run("query error", func(t *testing.T) {
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
	})

	t.Run("row scan error", func(t *testing.T) {
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
	})
}

func TestHasUserByPhone1(t *testing.T) {
	t.Parallel()

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

	t.Run("good query", func(t *testing.T) {
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
	})

	t.Run("query error", func(t *testing.T) {
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
	})

	t.Run("row scan error", func(t *testing.T) {
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
	})
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
//	repo := UserRepo{
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
