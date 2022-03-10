package validation

import (
	"regexp"
)

const phoneRegexp = `^[+]{1}7[(]{1}[0-9]{3}[)]{1}[0-9]{3}-[0-9]{2}-[0-9]{2}$`
const nameRegexp = `^[A-ZА-Я]{1}[a-zа-я]{1,}$`
const passwordRegexp = `^[A-Za-zА-Яа-я0-9]{8,}$`

var city = map[string]bool {
	"moscow": true,
	"voronezh": true,
}

func ValidatePhone(str string) bool {
	isMatch, _ := regexp.MatchString(phoneRegexp, str)
	return isMatch
}

func ValdateCity(str string) bool {
	_, cityExist := city[str]
	return cityExist
}

func ValidateUsername(str string) bool {
	isMatch, _ := regexp.MatchString(nameRegexp, str)
	return isMatch
}

func ValidatePassword(str string) bool {
	isMatch, _ := regexp.MatchString(passwordRegexp, str)
	return isMatch
}
