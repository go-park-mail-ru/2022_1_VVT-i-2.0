package ucase

import (
	"fmt"
	"regexp"
	"strings"

	addr "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
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

	toCutCityRegexpStr = `(?i)^ *гор\.|^ *гор |^ *г |^ *г\.|^ *город `
	separator          = ","
)

var (
	toCutCityRegexp = *regexp.MustCompile(toCutCityRegexpStr)
)

var defaultRes = []string{"Москва, ", "Москва, Бауманская Улица, ", "Москва, Тверская Улица, "}

type addressT struct {
	city       string
	street     streetT
	house      string
	toComplite int
}

func (u *AddrUcase) parseAddress(addrStr string) *addressT {
	addrParts := strings.Split(addrStr, separator)
	switch len(addrParts) - 1 {
	case 0:
		return &addressT{
			city:       strings.Trim(string(toCutCityRegexp.ReplaceAll([]byte(addrParts[0]), []byte(" "))), " "),
			toComplite: city,
		}
	case 1:
		return &addressT{
			city:       strings.Trim(string(toCutCityRegexp.ReplaceAll([]byte(addrParts[0]), []byte(" "))), " "),
			street:     *getStreet(addrParts[1]),
			toComplite: street,
		}

	default:
		return &addressT{
			city:       strings.Trim(string(toCutCityRegexp.ReplaceAll([]byte(addrParts[0]), []byte(" "))), " "),
			street:     *getStreet(addrParts[1]),
			house:      getHouse(addrParts[2]),
			toComplite: house,
		}
	}
}

func (u *AddrUcase) suggestCity() (*models.SuggestUcaseResp, error) {
	return &models.SuggestUcaseResp{Suggests: defaultRes}, nil
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
		fmt.Println(addr.city)
		return u.suggestCity()
	}
}
