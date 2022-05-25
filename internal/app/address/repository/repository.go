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

	if err != nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}

	suggsRepo := make([]string, len(suggs))
	for i, sugg := range suggs {
		suggsRepo[i] = *sugg
	}
	return &models.SuggestStreetRepoResp{StreetSuggests: suggsRepo}, nil
}

func (r *AddrRepo) SuggestHouse(address *models.SuggestHouseRepoReq) (*models.SuggestHouseRepoResp, error) {
	suggs := make([]*string, 0, address.SuggsLimit)
	err := r.DB.Select(&suggs, `SELECT house FROM houses WHERE street_id =$1 AND house ILIKE $2 LIMIT $3`, address.StreetId, address.House+"_%", address.SuggsLimit)

	if err != nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}

	suggsRepo := make([]string, len(suggs))
	for i, sugg := range suggs {
		suggsRepo[i] = *sugg
	}
	return &models.SuggestHouseRepoResp{HouseSuggests: suggsRepo}, nil
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

func (r *AddrRepo) GetTopUserAddrs(req *models.GetTopUserAddrsRepoReq) (*models.GetTopUserAddrsRepoResp, error) {
	addrs := make([]*string, 0, req.Limit)
	err := r.DB.Select(&addrs, `SELECT address FROM (SELECT (unnest(addresses)::order_address).address address, (unnest(addresses)::order_address).count count FROM users WHERE id=$1) AS addrs ORDER BY (count + (row_number() OVER() + 1)/2) DESC LIMIT $2`, req.UserId, req.Limit)

	if err != nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}

	addrsResp := &models.GetTopUserAddrsRepoResp{Addrs: make([]string, len(addrs))}
	for i, addr := range addrs {
		addrsResp.Addrs[i] = *addr
	}
	return addrsResp, nil
}

func (r *AddrRepo) SuggestUserAddrs(req *models.SuggestUserAddrsRepoReq) (*models.SuggestUserAddrsRepoResp, error) {
	addrs := make([]*string, 0, req.Limit)
	err := r.DB.Select(&addrs, `SELECT address FROM (SELECT (unnest(addresses)::order_address).address address, (unnest(addresses)::order_address).count count FROM users WHERE id=$1) as a  WHERE address ILIKE $2 AND address ILIKE $3 ORDER BY (count + (row_number() OVER())/2) DESC LIMIT $4`, req.UserId, req.Addr+"%", "%"+req.StreetType+"%", req.Limit)

	if err != nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}

	addrsResp := &models.SuggestUserAddrsRepoResp{Addrs: make([]string, len(addrs))}
	for i, addr := range addrs {
		addrsResp.Addrs[i] = *addr
	}
	return addrsResp, nil
}
