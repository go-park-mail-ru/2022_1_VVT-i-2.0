package ucase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	promocodes "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode"
	"github.com/pkg/errors"
)

type PromocodeUcase struct {
	PromoRepo promocodes.Repository
}

func NewPromocodeUcase(promocodeRepo promocodes.Repository) *PromocodeUcase {
	return &PromocodeUcase{
		PromoRepo: promocodeRepo,
	}
}

func (u *PromocodeUcase) GetAllPromocodes() (*models.GetPromocodesUcaseResp, error) {
	promosRepoResp, err := u.PromoRepo.GetAllPromocodes()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	promosResp := models.GetPromocodesUcaseResp{Promos: make([]models.PromocodeUcaseResp, len(promosRepoResp.Promos))}

	for i, promo := range promosRepoResp.Promos {
		promosResp.Promos[i] = models.PromocodeUcaseResp(promo)
	}

	return &promosResp, nil
}
