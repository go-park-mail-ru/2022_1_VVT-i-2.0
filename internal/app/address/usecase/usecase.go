package ucase

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"

	addr "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
)

type AddrUcase struct {
	AddrRepo      addr.Repository
	replaceRegexp regexp.Regexp
}

func NewAddrUcase(repo addr.Repository) *AddrUcase {
	return &AddrUcase{
		AddrRepo:      repo,
		replaceRegexp: *regexp.MustCompile(toDeleteRegexp),
	}
}

const (
	city = iota
	street
	house
	ready

	toDeleteRegexp = `^ *гор\.|^ *гор |^ *г |^ *г\.|^ *город | ул\.| ул | дом | д.| д `
	separator      = ","
)

var defaultRes = []string{"Москва, ", "Москва, Бауманская Улица, ", "Москва, Тверская Улица, "}

type addressT struct {
	city       string
	street     string
	house      string
	toComplite int
}

func (u *AddrUcase) parseAddress(addrStr string) *addressT {
	addrStr = string(u.replaceRegexp.ReplaceAll([]byte(addrStr), []byte(" ")))
	addrParts := strings.Split(addrStr, separator)
	for i := range addrParts {
		addrParts[i] = strings.TrimSpace(addrParts[i])
	}
	switch len(addrParts) - 1 {
	case 0:
		return &addressT{
			toComplite: city,
		}
	case 1:
		return &addressT{
			city:       addrParts[0],
			street:     addrParts[1],
			toComplite: street,
		}
	// case 2:
	// 	return &addressT{
	// 		city:       addrParts[0],
	// 		street:     addrParts[1],
	// 		toComplite: street,
	// 	}

	default:
		return &addressT{
			city:       addrParts[0],
			street:     addrParts[1],
			house:      addrParts[2],
			toComplite: house,
		}
	}
}

func (u *AddrUcase) suggestCity() (*models.SuggestUcaseResp, error) {
	return &models.SuggestUcaseResp{Suggests: defaultRes}, nil
}

func (u *AddrUcase) suggestStreet(address addressT) (*models.SuggestUcaseResp, error) {
	city, err := u.AddrRepo.GetCity(address.city)
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
	var suggs *models.SuggestStreetRepoAnsw

	pozToCut := len(address.street)
	for i := 0; i < 3 && suggs == nil && pozToCut >= 0; i++ {

		street := []rune(address.street[:(pozToCut)])
		suggs, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoInput{Street: string(street)})
		if len(address.street)-i*3 <= 0 {
			break
		}

		pozToCut--
		for j := 0; j < 2 && pozToCut >= 0; pozToCut-- {
			if utf8.ValidString(address.street[pozToCut:]) {
				j++
			}
		}
		pozToCut++
	}
	if suggs == nil {
		return nil, errors.Wrap(err, "error suggesting street")
	}
	var suggsResp models.SuggestUcaseResp
	for _, addr := range suggs.StreetSuggests {
		suggsResp.Suggests = append(suggsResp.Suggests, city.Name+separator+" "+addr+", ")
	}
	return &suggsResp, err
}

func (u *AddrUcase) suggestHouse(address addressT) (*models.SuggestUcaseResp, error) {
	city, err := u.AddrRepo.GetCity(address.city)
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
	street, err := u.AddrRepo.GetStreet(&models.GetStreetRepoInput{Street: address.street, CityId: city.CityId})
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

	house, err := u.AddrRepo.GetHouse(&models.GetHouseRepoInput{StreetId: street.StreetId, House: address.house})
	if err == nil && house != nil && house.House != "" {
		return &models.SuggestUcaseResp{Suggests: []string{city.Name + separator + " " + street.Name + ", " + house.House}, AddressFull: true}, nil
	}

	var suggs *models.SuggestHouseRepoAnsw
	pozToCut := len(address.house)
	for i := 0; i < 3 && suggs == nil && pozToCut >= 0; i++ {

		house := []rune(address.house[:pozToCut])
		// fmt.Println(string(house))
		suggs, err = u.AddrRepo.SuggestHouse(&models.SuggestHouseRepoInput{StreetId: street.StreetId, House: string(house)})
		if len(address.house)-i*1 <= 0 {
			break
		}

		for pozToCut--; !utf8.ValidString(address.house[pozToCut:]) && pozToCut >= 0; pozToCut-- {
		}
	}
	if suggs == nil {
		return nil, errors.Wrap(err, "error suggesting house")
	}
	var suggsResp models.SuggestUcaseResp
	for _, house := range suggs.HouseSuggests {
		suggsResp.Suggests = append(suggsResp.Suggests, city.Name+separator+" "+street.Name+", "+house)
	}
	suggsResp.AddressFull = true
	return &suggsResp, err
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
