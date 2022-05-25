package address

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Ucase interface {
	Suggest(req *models.SuggestUcaseReq) (*models.SuggestsUcaseResp, error)
}
