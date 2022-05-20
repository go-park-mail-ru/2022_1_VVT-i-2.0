package ucase

import (
	"strings"
	"unicode/utf8"

	addr "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/addrParser"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
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
	suggsLimit = 5
)

const (
	city = iota
	street
	house
	ready
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

var defaultRes = []string{"Москва, "}

func (u *AddrUcase) suggestCity() (*models.SuggestUcaseResp, error) {
	return &models.SuggestUcaseResp{Suggests: defaultRes}, nil
}

func (u *AddrUcase) suggestStreet(address addressT) (*models.SuggestUcaseResp, error) {
	city, err := u.AddrRepo.GetCity(&models.GetCityRepoReq{City: address.city})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		switch {
		case cause == nil:
			{
				return nil, errors.Wrap(err, "unexpected error")
			}
		case cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB:
			{
				return u.suggestCity()
			}
		default:
			return nil, errors.Wrap(err, "error getting city from db")
		}
	}

	var suggsResp models.SuggestUcaseResp
	completeMachStreet, err := u.AddrRepo.GetStreet(&models.GetStreetRepoReq{Street: address.street.Name, StreetType: address.street.StreetType, CityId: city.CityId})
	NoComrleteMachSuggsLimit := suggsLimit
	if err == nil && completeMachStreet != nil && completeMachStreet.Name != "" {
		suggsResp.Suggests = append(suggsResp.Suggests, addrParser.ConcatAddr(city.Name, completeMachStreet.Name, ""))
		NoComrleteMachSuggsLimit--
	}

	suggs := &models.SuggestStreetRepoResp{}
	suggsFromMiddle := &models.SuggestStreetRepoResp{} // на запрос "Парковая" -> ответ "1-я Парковая"

	pozToCut := len(address.street.Name)
	for i := 0; i < 4 && len(suggs.StreetSuggests) == 0 && pozToCut >= 0; i++ {
		street := strings.TrimRight(string(address.street.Name[:(pozToCut)]), " ")

		suggs, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoReq{Street: street, StreetType: address.street.StreetType, SearchInMiddle: false, SuggsLimit: NoComrleteMachSuggsLimit})
		if suggs == nil {
			suggs = &models.SuggestStreetRepoResp{}
		}

		if len(suggs.StreetSuggests) < suggsLimit {
			suggsFromMiddle, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoReq{Street: street, StreetType: address.street.StreetType, SearchInMiddle: true, SuggsLimit: NoComrleteMachSuggsLimit - len(suggs.StreetSuggests)})
			if suggsFromMiddle != nil {
				suggs.StreetSuggests = append(suggs.StreetSuggests, suggsFromMiddle.StreetSuggests...)
			}
		}

		pozToCut--
		for j := 0; j < letterCountToTrimStreetPerIter && pozToCut >= 0; pozToCut-- {
			if utf8.ValidString(address.street.Name[pozToCut:]) {
				j++
			}
		}
		pozToCut++
	}

	if (suggs == nil || len(suggsResp.Suggests) == 0) && len(suggsResp.Suggests) > 0 {
		return nil, errors.Wrap(err, "error suggesting house")
	}

	if suggs != nil {
		for _, street := range suggs.StreetSuggests {
			if completeMachStreet == nil || street != completeMachStreet.Name {
				suggsResp.Suggests = append(suggsResp.Suggests, addrParser.ConcatAddr(city.Name, street, ""))
			}
		}
	}

	return &suggsResp, nil
}

func (u *AddrUcase) suggestHouse(address addressT) (*models.SuggestUcaseResp, error) {
	city, err := u.AddrRepo.GetCity(&models.GetCityRepoReq{City: address.city})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		switch {
		case cause == nil:
			{
				return nil, errors.Wrap(err, "unexpected error")
			}
		case cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB:
			{
				return u.suggestCity()
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
				return u.suggestStreet(address)
			}
		default:
			return nil, errors.Wrap(err, "error getting street from db")
		}
	}

	completeMachHouse, err := u.AddrRepo.GetHouse(&models.GetHouseRepoReq{StreetId: street.StreetId, House: address.house})
	var suggsResp models.SuggestUcaseResp
	NoComrleteMachSuggsLimit := suggsLimit
	if err == nil && completeMachHouse != nil && completeMachHouse.House != "" {
		suggsResp.Suggests = append(suggsResp.Suggests, addrParser.ConcatAddr(city.Name, street.Name, completeMachHouse.House))
		NoComrleteMachSuggsLimit--

	}

	var suggs *models.SuggestHouseRepoResp
	pozToCut := len(address.house)
	for i := 0; i < 3 && suggs == nil && pozToCut >= 0; i++ {
		house := []rune(address.house[:pozToCut])
		suggs, err = u.AddrRepo.SuggestHouse(&models.SuggestHouseRepoReq{StreetId: street.StreetId, House: string(house), SuggsLimit: NoComrleteMachSuggsLimit})
		if len(house)-i*1 <= 0 {
			break
		}

		for pozToCut--; !utf8.ValidString(address.house[pozToCut:]) && pozToCut >= 0; pozToCut-- {
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
			suggsResp.Suggests = append(suggsResp.Suggests, addrParser.ConcatAddr(city.Name, street.Name, house))
		}
	}
	suggsResp.AddressFull = true
	return &suggsResp, nil
}

func (u *AddrUcase) Suggest(address *models.SuggestUcaseReq) (*models.SuggestUcaseResp, error) {
	if address == nil {
		return u.suggestCity()
	}
	addr := u.parseAddress(address.Address)
	if addr == nil {
		return u.suggestCity()
	}
	switch addr.toComplite {
	case street:
		{
			return u.suggestStreet(*addr)
		}
	case house:
		{
			return u.suggestHouse(*addr)
		}
	default:
		return u.suggestCity()
	}
}

// если доболнять город -- 1 саджест на город + 5 саджестов
// все остальное : 3 саджеста на мои адреса + 3 саджеста на другое
