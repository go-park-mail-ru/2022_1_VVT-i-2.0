package ucase

import (
	"strings"
	"unicode/utf8"

	addr "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/addrParser"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/pkg/errors"
)

type AddrUcase struct {
	AddrRepo addr.Repository
}

func NewAddrUcase(repo addr.Repository) *AddrUcase {
	return &AddrUcase{
		AddrRepo: repo,
	}
}

const (
	suggsLimit = 6
)

const (
	city = iota
	street
	house
	letterCountToTrimStreetPerIter = 2
)

type addressT struct {
	city       string
	street     addrParser.StreetT
	house      string
	toComplite int
}

func (u *AddrUcase) parseAddress(addrStr string) *addressT {
	addrParts := strings.Split(addrStr, addrParser.Separator)
	switch len(addrParts) - 1 {
	case 0:
		return &addressT{
			city:       addrParser.GetCity(addrParts[0]),
			toComplite: city,
		}
	case 1:
		return &addressT{
			city:       addrParser.GetCity(addrParts[0]),
			street:     *addrParser.GetStreet(addrParts[1]),
			toComplite: street,
		}

	default:
		return &addressT{
			city:       addrParser.GetCity(addrParts[0]),
			street:     *addrParser.GetStreet(addrParts[1]),
			house:      addrParser.GetHouse(addrParts[2]),
			toComplite: house,
		}
	}
}

var cityResp = []models.OneSuggestUcaseResp{{Address: "Москва, ", Full: false}}

func (u *AddrUcase) suggestCity(limit int) (*models.SuggestsUcaseResp, error) {
	return &models.SuggestsUcaseResp{Suggests: cityResp}, nil
}

func (u *AddrUcase) getTopUserAddrs(userId int64, limit int) (*models.SuggestsUcaseResp, error) {
	if !validator.IsUserId(userId) {
		return nil, nil
	}
	userAddrsRepo, err := u.AddrRepo.GetTopUserAddrs(&models.GetTopUserAddrsRepoReq{UserId: userId, Limit: limit - len(cityResp)})
	if err != nil {
		return nil, errors.Wrap(err, "error getting user address from storage")
	}
	resp := make([]models.OneSuggestUcaseResp, len(userAddrsRepo.Addrs))
	for i, addr := range userAddrsRepo.Addrs {
		resp[i] = models.OneSuggestUcaseResp{Address: addr, Full: true}
	}
	return &models.SuggestsUcaseResp{Suggests: resp}, nil
}

func (u *AddrUcase) suggestStreet(cityId int64, cityName string, streetReq addrParser.StreetT, limit int) (*models.SuggestsUcaseResp, error) {

	var suggsResp models.SuggestsUcaseResp
	completeMachStreet, err := u.AddrRepo.GetStreet(&models.GetStreetRepoReq{Street: streetReq.Name, StreetType: streetReq.StreetType, CityId: cityId})
	noComleteMachSuggsLimit := limit
	if err == nil && completeMachStreet != nil && completeMachStreet.Name != "" {
		suggsResp.Suggests = append(suggsResp.Suggests, models.OneSuggestUcaseResp{Address: addrParser.ConcatAddr(cityName, completeMachStreet.Name, ""), Full: false})
		noComleteMachSuggsLimit--
	}

	suggs := &models.SuggestStreetRepoResp{}
	var suggsFromMiddle *models.SuggestStreetRepoResp // на запрос "Парковая" -> ответ "1-я Парковая"

	pozToCut := len(streetReq.Name)
	for i := 0; i < 4 && len(suggs.StreetSuggests) == 0 && pozToCut >= 0; i++ {
		street := strings.TrimRight(string(streetReq.Name[:(pozToCut)]), " ")

		suggs, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoReq{Street: street, StreetType: streetReq.StreetType, SearchInMiddle: false, SuggsLimit: noComleteMachSuggsLimit})
		if suggs == nil {
			suggs = &models.SuggestStreetRepoResp{}
		}

		if len(suggs.StreetSuggests) < limit {
			suggsFromMiddle, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoReq{Street: street, StreetType: streetReq.StreetType, SearchInMiddle: true, SuggsLimit: noComleteMachSuggsLimit - len(suggs.StreetSuggests)})
			if suggsFromMiddle != nil {
				suggs.StreetSuggests = append(suggs.StreetSuggests, suggsFromMiddle.StreetSuggests...)
			}
		}

		pozToCut--
		for j := 0; j < letterCountToTrimStreetPerIter && pozToCut >= 0; pozToCut-- {
			if utf8.ValidString(streetReq.Name[pozToCut:]) {
				j++
			}
		}
		pozToCut++
	}

	if (suggs == nil || len(suggsResp.Suggests) == 0) && len(suggsResp.Suggests) > 0 {
		return nil, errors.Wrap(err, "error suggesting house")
	}

	if suggs != nil {
		uniqueSuggs := map[string]int{}
		for _, street := range suggs.StreetSuggests {
			if uniqueSuggs[street] == 0 && (completeMachStreet == nil || street != completeMachStreet.Name) {
				suggsResp.Suggests = append(suggsResp.Suggests, models.OneSuggestUcaseResp{Address: addrParser.ConcatAddr(cityName, street, ""), Full: false})
			}
			uniqueSuggs[street] = 1
		}
	}

	return &suggsResp, nil
}

func (u *AddrUcase) suggestHouse(cityId int64, cityName string, streetId int64, street string, houseReq string, limit int) (*models.SuggestsUcaseResp, error) {
	completeMachHouse, err := u.AddrRepo.GetHouse(&models.GetHouseRepoReq{StreetId: streetId, House: houseReq})
	var suggsResp models.SuggestsUcaseResp
	NoComrleteMachSuggsLimit := suggsLimit
	if err == nil && completeMachHouse != nil && completeMachHouse.House != "" {
		suggsResp.Suggests = append(suggsResp.Suggests, models.OneSuggestUcaseResp{Address: addrParser.ConcatAddr(cityName, street, completeMachHouse.House), Full: true})
		NoComrleteMachSuggsLimit--
	}

	suggs := &models.SuggestHouseRepoResp{}
	pozToCut := len(houseReq)
	for i := 0; i < 3 && (suggs == nil || len(suggs.HouseSuggests) == 0) && pozToCut >= 0; i++ {
		house := []rune(houseReq[:pozToCut])
		suggs, err = u.AddrRepo.SuggestHouse(&models.SuggestHouseRepoReq{StreetId: streetId, House: string(house), SuggsLimit: NoComrleteMachSuggsLimit})
		if len(house)-i*1 <= 0 {
			break
		}

		for pozToCut--; !utf8.ValidString(houseReq[pozToCut:]) && pozToCut >= 0; pozToCut-- {
		}
	}

	if (suggs == nil || len(suggsResp.Suggests) == 0) && len(suggsResp.Suggests) > 0 {
		return &suggsResp, nil
	}
	if suggs == nil {
		return &suggsResp, err
	}
	for _, house := range suggs.HouseSuggests {
		if completeMachHouse == nil || house != completeMachHouse.House {
			suggsResp.Suggests = append(suggsResp.Suggests, models.OneSuggestUcaseResp{Address: addrParser.ConcatAddr(cityName, street, house), Full: true})
		}
	}
	return &suggsResp, nil
}

func (u *AddrUcase) suggestStreetUserAddrs(userId int64, req string, streetType string, limit int) (*models.SuggestsUcaseResp, error) {
	if !validator.IsUserId(userId) {
		return nil, nil
	}
	userAddrsRepo, err := u.AddrRepo.SuggestUserAddrs(&models.SuggestUserAddrsRepoReq{Addr: req, UserId: userId, StreetType: streetType, Limit: limit})
	if err != nil {
		return nil, errors.Wrap(err, "error getting user addresses from storage")
	}
	resp := make([]models.OneSuggestUcaseResp, len(userAddrsRepo.Addrs))
	for i, addr := range userAddrsRepo.Addrs {
		resp[i] = models.OneSuggestUcaseResp{Address: addr, Full: true}
	}
	return &models.SuggestsUcaseResp{Suggests: resp}, nil
}

func (u *AddrUcase) suggestHouseUserAddrs(userId int64, req string, limit int) (*models.SuggestsUcaseResp, error) {
	if !validator.IsUserId(userId) {
		return nil, nil
	}
	userAddrsRepo, err := u.AddrRepo.SuggestUserAddrs(&models.SuggestUserAddrsRepoReq{Addr: req, UserId: userId, Limit: limit})
	if err != nil {
		return nil, errors.Wrap(err, "error getting user addresses from storage")
	}
	resp := make([]models.OneSuggestUcaseResp, len(userAddrsRepo.Addrs))
	for i, addr := range userAddrsRepo.Addrs {
		resp[i] = models.OneSuggestUcaseResp{Address: addr, Full: true}
	}
	return &models.SuggestsUcaseResp{Suggests: resp}, nil
}

func (u *AddrUcase) suggestCityMain(userId int64, address addressT, limit int) (*models.SuggestsUcaseResp, error) {
	topUserAddrs, _ := u.getTopUserAddrs(userId, limit-1)
	suggs, err := u.suggestCity(1)
	if err != nil {
		return nil, err
	}
	if topUserAddrs == nil {
		return &models.SuggestsUcaseResp{Suggests: suggs.Suggests}, nil
	}
	return &models.SuggestsUcaseResp{Suggests: append(topUserAddrs.Suggests, suggs.Suggests...)}, nil
}

func (u *AddrUcase) suggestStreetMain(userId int64, address addressT, limit int) (*models.SuggestsUcaseResp, error) {
	city, err := u.AddrRepo.GetCity(&models.GetCityRepoReq{City: address.city})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		switch {
		case cause == nil:
			{
				return nil, errors.Wrap(err, "unexpected error getting city")
			}
		case cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB:
			{
				return u.suggestCityMain(userId, address, limit)
			}
		default:
			return nil, errors.Wrap(err, "error getting city from db")
		}
	}

	suggs, err := u.suggestStreet(city.CityId, city.Name, address.street, limit)
	if err != nil {
		return nil, errors.Wrap(err, "error suggesting street")
	}

	userAddrs, err := u.suggestStreetUserAddrs(userId, addrParser.ConcatAddrToComplete(city.Name, address.street.Name, ""), address.street.StreetType, limit/2)
	if userAddrs == nil {
		return &models.SuggestsUcaseResp{Suggests: suggs.Suggests}, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "error suggesting street")
	}

	if limit-len(userAddrs.Suggests) >= len(suggs.Suggests) {
		return &models.SuggestsUcaseResp{Suggests: append(userAddrs.Suggests, suggs.Suggests...)}, nil
	}
	return &models.SuggestsUcaseResp{Suggests: append(userAddrs.Suggests, suggs.Suggests[:limit-len(userAddrs.Suggests)]...)}, nil

}

func (u *AddrUcase) suggestHouseMain(userId int64, address addressT, limit int) (*models.SuggestsUcaseResp, error) {
	city, err := u.AddrRepo.GetCity(&models.GetCityRepoReq{City: address.city})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		switch {
		case cause == nil:
			{
				return nil, errors.Wrap(err, "unexpected error getting city")
			}
		case cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB:
			{
				return u.suggestCityMain(userId, address, limit)
			}
		default:
			return nil, errors.Wrap(err, "error getting city from db")
		}
	}
	street, err := u.AddrRepo.GetStreet(&models.GetStreetRepoReq{Street: address.street.Name, StreetType: address.street.StreetType, CityId: city.CityId})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		switch {
		case cause == nil:
			{
				return nil, errors.Wrap(err, "unexpected error")
			}
		case cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB:
			{
				return u.suggestStreetMain(userId, address, limit)
			}
		default:
			return nil, errors.Wrap(err, "error getting street from db")
		}
	}

	suggs, err := u.suggestHouse(city.CityId, city.Name, street.StreetId, street.Name, address.house, limit)
	if err != nil {
		return nil, errors.Wrap(err, "error suggesting street")
	}

	userAddrs, err := u.suggestHouseUserAddrs(userId, addrParser.ConcatAddrToComplete(city.Name, street.Name, address.house), limit/2)
	if err != nil {
		return nil, errors.Wrap(err, "error suggesting street")
	}
	if userAddrs == nil {
		return &models.SuggestsUcaseResp{Suggests: suggs.Suggests}, nil
	}

	resp := models.SuggestsUcaseResp{Suggests: make([]models.OneSuggestUcaseResp, 0, limit)}
	uniqueSuggs := map[string]int{}
	for _, sugg := range append(userAddrs.Suggests, suggs.Suggests...) {
		if uniqueSuggs[sugg.Address] == 0 {
			resp.Suggests = append(resp.Suggests, sugg)
			uniqueSuggs[sugg.Address] = 1
		}
	}

	if len(resp.Suggests) < limit {
		return &models.SuggestsUcaseResp{Suggests: resp.Suggests}, nil
	}
	return &models.SuggestsUcaseResp{Suggests: resp.Suggests[:limit]}, nil

}

func (u *AddrUcase) Suggest(req *models.SuggestUcaseReq) (*models.SuggestsUcaseResp, error) {
	addr := u.parseAddress(req.Address)

	switch {
	case addr.toComplite == street:
		{
			return u.suggestStreetMain(req.UserId, *addr, suggsLimit)

		}
	case addr.toComplite == house:
		{
			return u.suggestHouseMain(req.UserId, *addr, suggsLimit)
		}
	}
	return u.suggestCityMain(req.UserId, *addr, suggsLimit)
}
