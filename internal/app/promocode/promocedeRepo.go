package restaurants

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetAllPromocodes() (*models.GetPromocodesRepoResp, error)
}
