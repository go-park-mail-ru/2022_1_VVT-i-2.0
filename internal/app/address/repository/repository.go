package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type AddrRepo struct {
	DB *sqlx.DB
}

func NewAddrRepo(db *sqlx.DB) *AddrRepo {
	return &AddrRepo{DB: db}
}

func (r *AddrRepo) SuggestStreet(address *models.SuggestStreetRepoReq) (*models.SuggestStreetRepoResp, error) {
	suggs := make([]*string, 0, address.SuggsLimit)

	var err error
	switch {
	case address.StreetType == "" && address.SearchInMiddle:
		err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 LIMIT $2`, "_%"+address.Street+"%", address.SuggsLimit)
	case address.StreetType == "" && !address.SearchInMiddle:
		err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 LIMIT $2`, address.Street+"%", address.SuggsLimit)
	case address.StreetType != "" && address.SearchInMiddle:
		err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 AND type LIKE $2 LIMIT $3`, "_%"+address.Street+"%", address.StreetType, address.SuggsLimit)
	default:
		err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 AND type LIKE $2 LIMIT $3`, address.Street+"%", address.StreetType, address.SuggsLimit)
	}

	switch err {
	case nil:
		{
			if len(suggs) == 0 {
				return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
			}
			suggsRepo := make([]string, len(suggs))
			for i := 0; i < address.SuggsLimit && i < len(suggs) && (suggs[i]) != nil; i++ {
				suggsRepo[i] = *suggs[i]
			}
			return &models.SuggestStreetRepoResp{StreetSuggests: suggsRepo}, nil
		}
	case sql.ErrNoRows:
		{
			return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
		}
	default:
		{
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
	}
}

func (r *AddrRepo) SuggestHouse(address *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error) {
	suggs := make([]*string, 0, address.SuggsLimit)
	err := r.DB.Select(&suggs, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2 LIMIT $3`, address.StreetId, address.House+"_%", address.SuggsLimit)

	switch err {
	case nil:
		if len(suggs) == 0 {
			return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
		}
		suggsRepo := make([]string, len(suggs))
		for i := 0; i < address.SuggsLimit && i < len(suggs) && (suggs[i]) != nil; i++ {
			suggsRepo[i] = *suggs[i]
		}
		return &models.SuggestHouseRepoResp{HouseSuggests: suggsRepo}, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AddrRepo) GetCity(city *models.GetCityRepoReq) (*models.GetCityRepoResp, error) {
	cityAnsw := &models.GetCityRepoResp{}
	err := r.DB.Get(cityAnsw, `SELECT id, name FROM cities WHERE name ILIKE $1`, city.City)

	switch err {
	case nil:
		return cityAnsw, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AddrRepo) GetStreet(street *models.GetStreetRepoReq) (*models.GetStreetRepoResp, error) {
	streetAnsw := &models.GetStreetRepoResp{}

	err := r.DB.Get(streetAnsw, `SELECT id, name FROM streets WHERE city_id = $1 AND main_name ILIKE $2 AND type ILIKE $3`, street.CityId, street.Street, street.StreetType)
	switch err {
	case nil:
		return streetAnsw, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AddrRepo) GetHouse(house *models.GetHouseRepoReq) (*models.GetHouseRepoResp, error) {
	houseAnsw := &models.GetHouseRepoResp{}
	err := r.DB.Get(houseAnsw, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2`, house.StreetId, house.House)
	switch err {
	case nil:
		return houseAnsw, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
