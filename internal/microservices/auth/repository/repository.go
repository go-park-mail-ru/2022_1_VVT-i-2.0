package repository

import (
	"database/sql"

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

func (r *AuthRepo) GetUserByPhone(phone models.UserByPhoneRepoReq) (*models.UserDataRepo, error) {
	user := &models.UserDataRepo{}
	err := r.DB.Get(user, `SELECT id, phone, email, name, avatar FROM users WHERE phone = $1`, phone.Phone)
	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AuthRepo) GetTopUserAddr(req *models.GetTopUserAddrRepoReq) (*models.GetTopUserAddrRepoResp, error) {
	addr := &models.GetTopUserAddrRepoResp{}
	err := r.DB.Get(addr, `SELECT address FROM (SELECT (unnest(addresses)::order_address).address address, (unnest(addresses)::order_address).count count FROM users WHERE id=$1) AS addrs ORDER BY (count + (row_number() OVER() + 1)/2) DESC LIMIT 1`, req.UserId)

	if err == sql.ErrConnDone || err == sql.ErrTxDone {
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
	return addr, nil
}

func (r *AuthRepo) AddUser(newUser *models.AddUserRepoReq) (*models.UserDataRepo, error) {
	user := &models.UserDataRepo{}
	err := r.DB.Get(user, `INSERT INTO users (name,phone,email) VALUES ($1,$2,$3) RETURNING id, name, phone, email`, newUser.Name, newUser.Phone, newUser.Email)

	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	return user, nil
}

func (r *AuthRepo) HasUserByPhone(phone models.UserByPhoneRepoReq) (models.HasSuchUserRepoResp, error) {
	user := &models.UserDataRepo{}
	err := r.DB.Get(user, `SELECT id FROM users WHERE phone = $1`, phone.Phone)
	switch err {
	case nil:
		return models.HasSuchUserRepoResp{IsRegistered: true}, nil
	case sql.ErrNoRows:
		return models.HasSuchUserRepoResp{IsRegistered: false}, nil
	default:
		return models.HasSuchUserRepoResp{IsRegistered: false}, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
