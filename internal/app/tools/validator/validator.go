package validator

import (
	"regexp"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

var (
	phoneRegexp   = regexp.MustCompile(`^7[94][0-9]{9}$`)
	codeRegexp    = regexp.MustCompile(`[0-9]{4}$`)
	nameRegexp    = regexp.MustCompile(`^[A-ZА-Я]{1}[a-zа-я]{2,25}$`)
	promoRegexp   = regexp.MustCompile(`^[A-ZА-Яa-zа-я0-9]{2,25}$`)
	slugRegexp    = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9\-]{1,128}$`)
	addressRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9 \,\.\/\-]{0,256}$`)
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
			code, ok := i.(string)
			if !ok {
				return false
			}
			return codeRegexp.MatchString(code)
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"slug",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			slug, ok := i.(string)
			if !ok {
				return false
			}
			return slugRegexp.MatchString(slug)
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"name",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			name, ok := i.(string)
			if !ok {
				return false
			}
			return nameRegexp.MatchString(name)
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"promocode",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			promocode, ok := i.(string)
			if !ok {
				return false
			}
			return promoRegexp.MatchString(promocode)
		}),
	)
	govalidator.CustomTypeTagMap.Set(
		"address",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			addr, ok := i.(string)
			if !ok {
				return false
			}
			return addressRegexp.MatchString(addr)
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
			return phoneRegexp.MatchString(phone)
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
	return slugRegexp.MatchString(str)
}

func IsUserId(num int64) bool {
	return num > 0
}
