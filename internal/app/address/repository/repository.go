package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

const (
	suggsLimit = 5
)

type AddrRepo struct {
	DB *sqlx.DB
}

func NewAddrRepo(db *sqlx.DB) *AddrRepo {
	return &AddrRepo{DB: db}
}

// user := &models.UserDataStorage{}
// 	err := r.DB.Get(user, `SELECT id, phone, email, name FROM users WHERE phone = $1`, phone)
// 	switch err {
// 	case nil:
// 		return user, nil
// 	case sql.ErrNoRows:
// 		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
// 	default:
// 		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
// 	}
// items := make([]*Item, 0, 10)
// err := repo.DB.Select(&items, "SELECT id, title, updated FROM items")

func (r *AddrRepo) SuggestStreet(address *models.SuggestStreetRepoInput) (*models.SuggestStreetRepoAnsw, error) {
	// suggs := make([]*[]rune, 0, suggsLimit)
	suggs := make([]*string, 0, suggsLimit)
	// err := repo.DB.Select(&items, "SELECT id, title, updated FROM items")
	// var suggs []string
	err := r.DB.Select(&suggs, `SELECT name FROM streets WHERE name ILIKE $1 LIMIT $2`, address.Street+"%", suggsLimit)
	// err := r.DB.Select(&suggs, `SELECT name FROM msc_streets WHERE name ILIKE $1 LIMIT $2`, "%"+"Измайл"+"%", suggsLimit)
	fmt.Println(suggs)
	fmt.Println(err)
	fmt.Println(len(suggs))
	switch err {
	case nil:
		if len(suggs) == 0 {
			return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
		}
		suggsRepo := make([]string, len(suggs))
		for i := 0; i < suggsLimit && i < len(suggs) && (suggs[i]) != nil; i++ {
			suggsRepo[i] = *suggs[i]
		}
		return &models.SuggestStreetRepoAnsw{StreetSuggests: suggsRepo}, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *AddrRepo) SuggestHouse(address *models.SuggestHouseRepoInput) (*models.SuggestHouseRepoAnsw, error) {
	suggs := make([]*string, 0, suggsLimit)

	err := r.DB.Select(&suggs, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2  LIMIT $3`, address.StreetId, address.House+"%", suggsLimit)
	fmt.Println(suggs)
	fmt.Println(err)
	fmt.Println(len(suggs))
	switch err {
	case nil:
		if len(suggs) == 0 {
			return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, "")
		}
		suggsRepo := make([]string, len(suggs))
		for i := 0; i < suggsLimit && i < len(suggs) && (suggs[i]) != nil; i++ {
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
	err := r.DB.Get(streetAnsw, `SELECT id as streetid, name FROM streets WHERE name ILIKE $1`, street.Street)
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
