package addrParser

import (
	"regexp"
	"strings"
)

const Separator = ","

var toCutCityRegexp = *regexp.MustCompile(`(?i)^ *гор\.|^ *гор |^ *г |^ *г\.|^ *город `)

var regexpStreetType = map[string]regexp.Regexp{
	"Улица":      *regexp.MustCompile(`(?i) ул\.| ул |ул\.|улица`),
	"Проспект":   *regexp.MustCompile(`(?i) пр-т|пр-т |проспект`),
	"Проезд":     *regexp.MustCompile(`(?i)пр-д|проезд`),
	"Площадь":    *regexp.MustCompile(`(?i)пл\.| пл |пл-дь|площадь`),
	"Переулок":   *regexp.MustCompile(`(?i)переулок| пер\.|пер |пер-к`),
	"Шоссе":      *regexp.MustCompile(`(?i)ш\.| ш |шоссе`),
	"Бульвар":    *regexp.MustCompile(`(?i)бульвар| бул\.|б-р `),
	"Набережная": *regexp.MustCompile(`(?i)набережная| наб\.| наб |н-я `),
	"Аллея":      *regexp.MustCompile(`(?i)аллея|а-я `),
	"Квартал":    *regexp.MustCompile(`(?i)квартал|к-л |кварт\.|кварт `),
	"Тупик":      *regexp.MustCompile(`(?i)тупик|т-к |туп\.|туп `),
	"Линия":      *regexp.MustCompile(`(?i)линия|л-я |лин\.`),
	"Просек":     *regexp.MustCompile(`(?i)просек|пр-к |прос\.|прос `),
	"Километр":   *regexp.MustCompile(`(?i)километр|км\.|кил\.|кил |км `),
}

var toCutDomRegexp = *regexp.MustCompile(`(?i) дом\.| д |д\.|дом `)
var cutExtraSpaceRegexp = regexp.MustCompile(`\s+`)

type StreetT struct {
	StreetType string
	Name       string
}

func GetCity(city string) string {
	city = cutExtraSpaceRegexp.ReplaceAllString(city, " ")
	return strings.Trim(string(toCutCityRegexp.ReplaceAll([]byte(city), []byte(" "))), " ")
}

func GetStreet(streetStr string) *StreetT {
	streetStr = cutExtraSpaceRegexp.ReplaceAllString(streetStr, " ")
	for streetType, typeRegexp := range regexpStreetType {
		cutStreet := string(typeRegexp.ReplaceAll([]byte(streetStr), []byte(" ")))
		if len(cutStreet) != len(streetStr) {
			return &StreetT{
				Name:       strings.Trim(cutStreet, " "),
				StreetType: streetType,
			}
		}
	}

	return &StreetT{
		Name:       strings.Trim(streetStr, " "),
		StreetType: "",
	}
}

func GetHouse(house string) string {
	house = cutExtraSpaceRegexp.ReplaceAllString(house, " ")
	return strings.Trim(string(toCutDomRegexp.ReplaceAll([]byte(house), []byte(" "))), " ")
}

func ConcatAddr(city string, street string, house string) string {
	return city + Separator + " " + street + Separator + " " + house
}
