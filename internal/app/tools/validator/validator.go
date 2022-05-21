package validator

import (
	"regexp"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

const (
	phoneRegexp   = `^7[94][0-9]{9}$`
	nameRegexp    = `^[A-ZА-Я]{1}[a-zа-я]{2,25}$`
	slugRegexp    = `^[a-zA-Zа-яА-Я0-9\-]{1,128}$`
	addressRegexp = `^[a-zA-Zа-яА-Я0-9 \,\.\/\-]{0,256}$`
	commentMaxLen = 1024
)

func init() {
	govalidator.CustomTypeTagMap.Set(
		"stars",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			stars, ok := i.(int)
			return ok && (stars <= 5) && (stars >= 1)
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"code",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			code, ok := i.(int)
			return ok && ((code <= 9999) && (code >= 0))
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"slug",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			slug, ok := i.(string)
			if !ok {
				return false
			}

			isSlug, _ := regexp.MatchString(slugRegexp, slug)
			return isSlug
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"name",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			name, ok := i.(string)
			if !ok {
				return false
			}

			isName, _ := regexp.MatchString(nameRegexp, name)
			return isName
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"address",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			addr, ok := i.(string)
			if !ok {
				return false
			}

			isAddr, _ := regexp.MatchString(addressRegexp, addr)
			return isAddr
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"comment",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			comment, ok := i.(string)
			return ok && len(comment) <= commentMaxLen && len(comment) > 0
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"phone",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			phone, ok := i.(string)
			if !ok {
				return false
			}

			isPhone, _ := regexp.MatchString(phoneRegexp, phone)
			return isPhone
		}),
	)

	govalidator.CustomTypeTagMap.Set(
		"userId",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			id, ok := i.(models.UserId)
			return ok && id > 0
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"expired",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			exp, ok := i.(time.Time)
			return ok && !exp.Before(time.Now())
		}),
	)
}

func IsSlug(str string) bool {
	isSlug, _ := regexp.MatchString(slugRegexp, str)
	return isSlug
}
