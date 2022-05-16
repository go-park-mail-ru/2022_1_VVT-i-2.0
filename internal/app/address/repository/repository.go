package repository

import (
	"database/sql"
	"strings"

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

func (r *AddrRepo) SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error) {
	suggs := make([]*string, 0, address.SuggsLimit)

	var err error
	switch {
	case address.StreetType == "" && address.SearchInMiddle:
		{
			err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 LIMIT $2`, "_%"+address.Street+"%", address.SuggsLimit)
		}
	case address.StreetType == "" && !address.SearchInMiddle:
		{
			err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 LIMIT $2`, address.Street+"%", address.SuggsLimit)
		}
	case address.StreetType != "" && address.SearchInMiddle:
		{
			err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 AND type LIKE $2 LIMIT $3`, "_%"+address.Street+"%", address.StreetType, address.SuggsLimit)
		}
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
			return &models.SuggestStreetRepoAnsw{StreetSuggests: suggsRepo}, nil
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

func (r *AddrRepo) SuggestHouse(address *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error) {
	suggs := make([]*string, 0, address.SuggsLimit)
	err := r.DB.Select(&suggs, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2  LIMIT $3`, address.StreetId, address.House+"%", address.SuggsLimit)

	if len(suggs) < address.SuggsLimit {
		suggs2 := make([]*string, 0, address.SuggsLimit)
		err = r.DB.Select(&suggs2, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2  LIMIT $3`, address.StreetId, "%"+address.House+"%", address.SuggsLimit-len(suggs))
		suggs = append(suggs, suggs2...)
	}
	switch err {
	case nil:
		if len(suggs) == 0 {
			return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
		}
		suggsRepo := make([]string, len(suggs))
		for i := 0; i < address.SuggsLimit && i < len(suggs) && (suggs[i]) != nil; i++ {
			suggsRepo[i] = *suggs[i]
		}
		return &models.SuggestHouseRepoAnsw{HouseSuggests: suggsRepo}, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AddrRepo) GetCity(city string) (*models.GetCityRepoAnsw, error) {
	if strings.ToLower(city) == "москва" {
		return &models.GetCityRepoAnsw{CityId: 0, Name: "Москва"}, nil
	}
	return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
}

func (r *AddrRepo) GetStreet(street *models.GetStreetRepoInput) (*models.GetStreetRepoAnsw, error) {
	streetAnsw := &models.GetStreetRepoAnsw{}

	// err = r.DB.Select(&suggs, `SELECT name FROM streets WHERE main_name ILIKE $1 AND type LIKE $2 LIMIT $3`, address.Street+"%", address.StreetType, address.SuggsLimit)
	// err := r.DB.Get(streetAnsw, `SELECT id as streetid, name FROM streets WHERE name ILIKE $1`, street.Street)
	// TODO надо удалять заменить несколько пробелов на один
	err := r.DB.Get(streetAnsw, `SELECT id, name FROM streets WHERE main_name ILIKE $1 AND type LIKE $2 LIMIT $3`, street.Street)
	switch err {
	case nil:
		return streetAnsw, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AddrRepo) GetHouse(house *models.GetHouseRepoInput) (*models.GetHouseRepoAnsw, error) {
	houseAnsw := &models.GetHouseRepoAnsw{}
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
