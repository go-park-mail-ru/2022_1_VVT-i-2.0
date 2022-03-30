package validator

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

func init() {
	govalidator.CustomTypeTagMap.Set(
		"userId",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			fmt.Println("==+=+===userid validate==+=+==")
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
			fmt.Println("-------------expires validate------------")
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
