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

func (r *UserRepo) GetUserById(id models.UserId) (*models.UserDataRepo, error) {
	user := &models.UserDataRepo{}
	err := r.DB.Get(user, `SELECT id, phone, email, name, avatar FROM users WHERE id = $1`, id)

	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *UserRepo) UpdateUser(updUser *models.UpdateUserStorage) (*models.UserDataRepo, error) {
	var err error
	user := &models.UserDataRepo{}

	switch {
	case updUser.Email != "" && updUser.Name != "" && updUser.Avatar != "":
		err = r.DB.Get(user, `UPDATE users SET name=$1, email=$2, avatar=$3 WHERE id=$4 RETURNING id, name, email, phone, avatar`, updUser.Name, updUser.Email, updUser.Avatar, updUser.Id)
	case updUser.Email != "" && updUser.Name == "" && updUser.Avatar == "":
		err = r.DB.Get(user, `UPDATE users SET email=$1 WHERE id=$2 RETURNING id, name, email, phone, avatar`, updUser.Email, updUser.Id)
	case updUser.Email == "" && updUser.Name != "" && updUser.Avatar == "":
		err = r.DB.Get(user, `UPDATE users SET name=$1 WHERE id=$2 RETURNING id, name, email, phone, avatar`, updUser.Name, updUser.Id)
	case updUser.Email == "" && updUser.Name == "" && updUser.Avatar != "":
		err = r.DB.Get(user, `UPDATE users SET avatar=$1 WHERE id=$2 RETURNING id, name, email, phone, avatar`, updUser.Avatar, updUser.Id)
	case updUser.Email != "" && updUser.Name != "" && updUser.Avatar == "":
		err = r.DB.Get(user, `UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING id, name, email, phone, avatar`, updUser.Name, updUser.Email, updUser.Id)
	case updUser.Email != "" && updUser.Name == "" && updUser.Avatar != "":
		err = r.DB.Get(user, `UPDATE users SET email=$1, avatar=$2 WHERE id=$3 RETURNING id, name, email, phone, avatar`, updUser.Email, updUser.Avatar, updUser.Id)
	case updUser.Email == "" && updUser.Name != "" && updUser.Avatar != "":
		err = r.DB.Get(user, `UPDATE users SET name=$1, avatar=$2 WHERE id=$3 RETURNING id, name, email, phone, avatar`, updUser.Name, updUser.Avatar, updUser.Id)
	default:
		return nil, nil
	}

	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		if err == sql.ErrNoRows {
			return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_UPDATE, err.Error())
	}
	return user, nil
}
