package auth

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type Repository interface {
	AddUser(newUser *models.AddUserRepoReq) (*models.UserDataRepo, error)
	HasUserByPhone(phone models.UserByPhoneRepoReq) (models.HasSuchUserRepoResp, error)
	GetUserByPhone(phone models.UserByPhoneRepoReq) (*models.UserDataRepo, error)
	GetTopUserAddr(req *models.GetTopUserAddrRepoReq) (*models.GetTopUserAddrRepoResp, error)
}
