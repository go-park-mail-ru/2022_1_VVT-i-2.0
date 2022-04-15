package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// case sql.ErrNoRows:
// 	return servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, (*err).Error())
// 	// return nil
// default:
// 	return servErrors.NewError(servErrors.DB_ERROR, (*err).Error())

func (r *UserRepo) GetUserByPhone(phone string) (*models.UserDataStorage, error) {
	user := &models.UserDataStorage{}
	err := r.DB.Get(user, `SELECT id, phone, email, name FROM users WHERE phone = $1`, phone)
	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *UserRepo) AddUser(newUser *models.UserAddDataStorage) (models.UserId, error) {
	result, err := r.DB.Exec(`INSERT INTO users (name,phone,email) VALUES ($1,$2,$3) RETURNING id`, newUser.Name, newUser.Phone, newUser.Email)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return 0, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return 0, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return 0, servErrors.NewError(servErrors.DB_INSERT, "")
	}
	id, _ := result.LastInsertId()
	return models.UserId(id), nil
}

func (r *UserRepo) GetUserById(id models.UserId) (*models.UserDataStorage, error) {
	user := &models.UserDataStorage{}
	err := r.DB.Get(user, `SELECT id, phone, email, name FROM users WHERE id = $1`, id)

	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *UserRepo) UpdateUser(updUser *models.UpdateUser) (*models.UserDataStorage, error) {
	var err error
	var result sql.Result
	switch {
	case updUser.Email != "" && updUser.Name != "":
		result, err = r.DB.Exec(`UPDATE users SET name=$1, email=$2 WHERE id=$3`, updUser.Name, updUser.Email, updUser.Id)

	case updUser.Email != "" && updUser.Name == "":
		result, err = r.DB.Exec(`UPDATE users SET email=$1 WHERE id=$2`, updUser.Email, updUser.Id)

	case updUser.Email == "" && updUser.Name != "":
		result, err = r.DB.Exec(`UPDATE users SET name=$1 WHERE id=$2`, updUser.Name, updUser.Id)
	default:
		return nil, nil
	}

	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_UPDATE, err.Error())
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return nil, servErrors.NewError(servErrors.DB_UPDATE, "")
	}
	return r.GetUserById(updUser.Id)
}

func (r *UserRepo) HasUserByPhone(phone string) (bool, error) {
	user := &models.UserDataStorage{}
	err := r.DB.Get(user, `SELECT id, phone, email, name FROM users WHERE phone = $1`, phone)
	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

/*
func (repo *RepoSqlx) GetAll() ([]*Item, error) {
	items := make([]*Item, 0, 10)
	err := repo.DB.Select(&items, "SELECT id, title, updated FROM items")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *RepoSqlx) GetAll_0() ([]*Item, error) {
	items := make([]*Item, 0, 10)
	rows, err := repo.DB.Queryx("SELECT id, title, updated FROM items")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := &Item{}
		// MapScan, SliceScan
		err := rows.StructScan(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (repo *RepoSqlx) GetByID(id int64) (*Item, error) {
	post := &Item{}
	err := repo.DB.Get(post, `SELECT id, title, updated, description FROM items WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *RepoSqlx) Add(elem *Item) (int64, error) {
	result, err := repo.DB.NamedExec(
		`INSERT INTO person (first_name,last_name,email) VALUES (:title, :description)`,
		map[string]interface{}{
			"title":       elem.Title,
			"description": elem.Description,
		})
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *RepoSqlx) Update(elem *Item) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE items SET"+
			"`title` = ?"+
			",`description` = ?"+
			",`updated` = ?"+
			"WHERE id = ?",
		elem.Title,
		elem.Description,
		"rvasily",
		elem.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (repo *RepoSqlx) Delete(id int64) (int64, error) {
	result, err := repo.DB.Exec(
		"DELETE FROM items WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
*/