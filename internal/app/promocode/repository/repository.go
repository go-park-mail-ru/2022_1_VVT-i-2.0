package repository

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type PromocodeRepo struct {
	DB *sqlx.DB
}

func NewPromocodeRepo(db *sqlx.DB) *PromocodeRepo {
	return &PromocodeRepo{DB: db}
}

func (r *PromocodeRepo) GetAllPromocodes() (*models.GetPromocodesRepoResp, error) {
	promocodes := make([]*models.PromocodeRepoResp, 0)
	err := r.DB.Select(&promocodes, `SELECT r.name, r.slug, p.text, p.discount, p.price_reduction, p.min_price, p.promocode, p.image_path, p.logo_path FROM promocodes p JOIN restaurants r ON p.restaurant_id=r.id`)
	switch err {
	case nil:
		resp := models.GetPromocodesRepoResp{Promos: make([]models.PromocodeRepoResp, len(promocodes))}
		for i, promocode := range promocodes {
			resp.Promos[i] = models.PromocodeRepoResp(*promocode)
		}
		return &resp, nil
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
