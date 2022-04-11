package validator

import (
	"regexp"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

const (
	phoneRegexp = `^79[0-9]{9}$`
)

func init() {
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
			if !ok {
				return false
			}
			if id <= 0 {
				return false
			}
			return true
		}),
	)

	govalidator.CustomTypeTagMap.Set(
		"expired",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			exp, ok := i.(time.Time)
			if !ok {
				return false
			}

			if exp.Before(time.Now()) {
				return false
			}
			return true
		}),
	)
}
