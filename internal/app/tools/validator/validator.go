package validator

import (
	"regexp"
	"time"
)

var (
	phoneRegexp        = regexp.MustCompile(`^7[94][0-9]{9}$`)
	codeRegexp         = regexp.MustCompile(`[0-9]{4}$`)
	nameRegexp         = regexp.MustCompile(`^[A-ZА-ЯЁ]{1}[a-zа-яё]{2,25}$`)
	promoRegexp        = regexp.MustCompile(`^[A-ZА-Яa-zа-я0-9]{2,25}$`)
	slugRegexp         = regexp.MustCompile(`^[a-zA-Z0-9\-]{1,128}$`)
	addressRegexp      = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ0-9 \,\.\/\-]{0,256}$`)
	letterNumberRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ0-9 \,\.\/\-]{0,25}$`)
	categoryRegexp     = regexp.MustCompile(`^[a-zA-Zа-яАёЁ-Я0-9 \,\.\/\-]{2,100}$`)
	emailRegexp        = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	commentMaxLen      = 1024
	reviewMaxLen       = 1024
	restaurantMaxLen   = 100
	intercomMaxLen     = 25
)

func IsPhone(str string) bool {
	return phoneRegexp.MatchString(str)
}

func IsAuthCode(str string) bool {
	return codeRegexp.MatchString(str)
}

func IsName(str string) bool {
	return nameRegexp.MatchString(str)
}

func IsSlug(str string) bool {
	return slugRegexp.MatchString(str)
}

func IsUserId(num int64) bool {
	return num > 0
}

func IsOrderId(num int64) bool {
	return num > 0
}

func IsStars(num int) bool {
	return (num <= 5) && (num >= 1)
}

func IsAddress(str string) bool {
	return addressRegexp.MatchString(str)
}

func IsPromocode(str string) bool {
	return promoRegexp.MatchString(str)
}

func IsComment(str string) bool {
	return len(str) <= commentMaxLen
}

func IsReview(str string) bool {
	return len(str) <= reviewMaxLen
}

func IsNotExpired(date time.Time) bool {
	return !date.Before(time.Now())
}

func IsEntrance(str string) bool {
	return letterNumberRegexp.MatchString(str)
}

func IsIntercom(str string) bool {
	return len(str) <= intercomMaxLen
}

func IsFloor(str string) bool {
	return letterNumberRegexp.MatchString(str)
}

func IsFlat(str string) bool {
	return letterNumberRegexp.MatchString(str)
}

func IsRestaurantCategory(str string) bool {
	return categoryRegexp.MatchString(str)
}

func IsRestaurantQuery(str string) bool {
	return len(str) <= restaurantMaxLen
}

func IsEmail(str string) bool {
	return emailRegexp.MatchString(str)
}
