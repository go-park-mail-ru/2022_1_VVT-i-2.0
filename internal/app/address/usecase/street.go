package ucase

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/pkg/errors"
)

const (
	toCutUlitsaRegexpStr      = `(?i) ул\.| ул |ул\.|улица`
	toCutProspectRegexpStr    = `(?i) пр-т|пр-т |проспект`
	toCutProezdRegexpStr      = `(?i)пр-д|проезд`
	toCutPloshadRegexpStr     = `(?i)пл\.| пл |пл-дь|площадь`
	toCutPereulokRegexpStr    = `(?i)переулок| пер\.|пер |пер-к`
	toCutShosseRegexpStr      = `(?i)ш\.| ш |шоссе`
	toCutBulvarRegexpStr      = `(?i)бульвар| бул\.|б-р `
	toCutNabereznayaRegexpStr = `(?i)набережная| наб\.| наб |н-я `
	toCutAlleyaRegexpStr      = `(?i)аллея|а-я `
	toCutKvartalRegexpStr     = `(?i)квартал|к-л |кварт\.|кварт `
	toCutTupikRegexpStr       = `(?i)тупик|т-к |туп\.|туп `
	toCutLinyaRegexpStr       = `(?i)линия|л-я |лин\.`
	toCutProsekRegexpStr      = `(?i)просек|пр-к |прос\.|прос `
	toCutKilometrRegexpStr    = `(?i)километр|км\.|кил\.|кил |км `
)

const (
	ulitsa      = "Улица"
	prospect    = "Проспект"
	proezd      = "Проезд"
	bulvar      = "Бульвар"
	shosse      = "Шоссе"
	ploshad     = "Площадь"
	kvartal     = "Квартал"
	pereulok    = "Переулок"
	nabereznaya = "Набережная"
	alleya      = "Аллея"
	tupik       = "Тупик"
	linya       = "Линия"
	prosek      = "Просек"
	kilometr    = "Километр"
)

var regexpStreetType = map[string]regexp.Regexp{
	ulitsa:      *regexp.MustCompile(toCutUlitsaRegexpStr),
	prospect:    *regexp.MustCompile(toCutProspectRegexpStr),
	proezd:      *regexp.MustCompile(toCutProezdRegexpStr),
	ploshad:     *regexp.MustCompile(toCutPloshadRegexpStr),
	pereulok:    *regexp.MustCompile(toCutPereulokRegexpStr),
	shosse:      *regexp.MustCompile(toCutShosseRegexpStr),
	bulvar:      *regexp.MustCompile(toCutBulvarRegexpStr),
	nabereznaya: *regexp.MustCompile(toCutNabereznayaRegexpStr),
	alleya:      *regexp.MustCompile(toCutAlleyaRegexpStr),
	kvartal:     *regexp.MustCompile(toCutKvartalRegexpStr),
	tupik:       *regexp.MustCompile(toCutTupikRegexpStr),
	linya:       *regexp.MustCompile(toCutLinyaRegexpStr),
	prosek:      *regexp.MustCompile(toCutProsekRegexpStr),
	kilometr:    *regexp.MustCompile(toCutKilometrRegexpStr),
}

const (
	letterCountToTrim = 2
)

type streetT struct {
	streetType string
	name       string
}

func getStreet(streetStr string) *streetT {
	for streetType, typeRegexp := range regexpStreetType {
		cutStreet := string(typeRegexp.ReplaceAll([]byte(streetStr), []byte(" ")))
		if len(cutStreet) != len(streetStr) {
			return &streetT{
				name:       strings.Trim(cutStreet, " "),
				streetType: streetType,
			}
		}
	}

	return &streetT{
		name:       strings.Trim(streetStr, " "),
		streetType: "",
	}
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

	var suggsResp models.SuggestUcaseResp
	street, err := u.AddrRepo.GetStreet(&models.GetStreetRepoInput{Street: address.street.name, StreetType: address.street.streetType, CityId: city.CityId})
	if err == nil && street != nil && street.Name != "" {
		suggsResp.Suggests = append(suggsResp.Suggests, city.Name+separator+" "+street.Name+", ")
	}

	suggs := &models.SuggestStreetRepoAnsw{}
	suggsFromMiddle := &models.SuggestStreetRepoAnsw{} // на запрос "Парковая" -> ответ "1-я Парковая"

	pozToCut := len(address.street.name)
	for i := 0; i < 4 && len(suggs.StreetSuggests) == 0 && pozToCut >= 0; i++ {
		street := strings.TrimRight(string(address.street.name[:(pozToCut)]), " ")

		suggs, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoInput{Street: street, StreetType: address.street.streetType, SearchInMiddle: false, SuggsLimit: suggsLimit})
		if suggs == nil {
			suggs = &models.SuggestStreetRepoAnsw{}
		}
		if len(suggs.StreetSuggests) < suggsLimit {
			suggsFromMiddle, err = u.AddrRepo.SuggestStreet(&models.SuggestStreetRepoInput{Street: street, StreetType: address.street.streetType, SearchInMiddle: true, SuggsLimit: suggsLimit - len(suggs.StreetSuggests)})
			if suggsFromMiddle != nil {
				suggs.StreetSuggests = append(suggs.StreetSuggests, suggsFromMiddle.StreetSuggests...)
			}
		}

		pozToCut--
		for j := 0; j < letterCountToTrim && pozToCut >= 0; pozToCut-- {
			if utf8.ValidString(address.street.name[pozToCut:]) {
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
			suggsResp.Suggests = append(suggsResp.Suggests, city.Name+separator+" "+street+", ")
		}
	}

	return &suggsResp, nil
}
