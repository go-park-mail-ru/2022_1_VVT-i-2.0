package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	DB *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

// func (r *UserRepo) GetUserByPhone(phone string) (*models.UserDataStorage, error) {
// 	user := &models.UserDataStorage{}
// 	err := r.DB.Get(user, `SELECT id, phone, email, name, avatar FROM users WHERE phone = $1`, phone)
// 	fmt.Println(err)
// 	fmt.Println(user)
// 	switch err {
// 	case nil:
// 		return user, nil
// 	case sql.ErrNoRows:
// 		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
// 	default:
// 		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
// 	}
// }

// func (r *UserRepo) AddUser(newUser *models.UserAddDataStorage) (*models.UserDataStorage, error) {
// 	user := &models.UserDataStorage{}
// 	err := r.DB.Get(user, `INSERT INTO users (name,phone,email) VALUES ($1,$2,$3) RETURNING id, name, phone, email`, newUser.Name, newUser.Phone, newUser.Email)

// 	if err != nil {
// 		if err == sql.ErrConnDone || err == sql.ErrTxDone {
// 			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
// 		}
// 		return nil, servErrors.NewError(servErrors.DB_INSERT, err.Error())
// 	}
// 	if user == nil {
// 		return nil, servErrors.NewError(servErrors.DB_INSERT, "")
// 	}
// 	return user, nil
// }

// func (r *UserRepo) GetUserById(id models.UserId) (*models.UserDataStorage, error) {
// 	user := &models.UserDataStorage{}
// 	err := r.DB.Get(user, `SELECT id, phone, email, name, avatar FROM users WHERE id = $1`, id)

// 	switch err {
// 	case nil:
// 		return user, nil
// 	case sql.ErrNoRows:
// 		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
// 	default:
// 		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
// 	}
// }

func (r *AuthRepo) HasUserByPhone(phone models.SendCodeRepoReq) (models.SendCodeRepoResp, error) {
	user := &models.UserDataStorage{}
	err := r.DB.Get(user, `SELECT id FROM users WHERE phone = $1`, phone.Phone)
	fmt.Println(err)
	fmt.Println(user)
	switch err {
	case nil:
		return models.SendCodeRepoResp{IsRegistered: true}, nil
	case sql.ErrNoRows:
		return models.SendCodeRepoResp{IsRegistered: false}, nil
	default:
		return models.SendCodeRepoResp{IsRegistered: false}, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
