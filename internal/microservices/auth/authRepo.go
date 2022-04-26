package auth

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type Repository interface {
	// AddUser(newUser *models.UserAddDataStorage) (*models.UserDataStorage, error)
	// GetUserByPhone(phone string) (*models.UserDataStorage, error)
	// GetUserById(id models.UserId) (*models.UserDataStorage, error)
	HasUserByPhone(phone models.SendCodeRepoReq) (models.SendCodeRepoResp, error)
}
