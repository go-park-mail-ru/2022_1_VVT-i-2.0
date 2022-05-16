package ucase

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
)

const (
	toCutDomRegexpStr = `(?i) дом\.| д |д\.|дом `
)

var toCutDomRegexp = *regexp.MustCompile(toCutDomRegexpStr)

func getHouse(house string) string {
	return strings.Trim(string(toCutDomRegexp.ReplaceAll([]byte(house), []byte(" "))), " ")
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
	street, err := u.AddrRepo.GetStreet(&models.GetStreetRepoInput{Street: address.street.name + " " + address.street.streetType, CityId: city.CityId})
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
