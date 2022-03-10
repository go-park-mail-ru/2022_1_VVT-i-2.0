package validation

import "regexp"

var city = map[string]bool {
	"moscow": true,
	"voronezh": true,
}

func ValidatePhone(str string) bool {
	isMatch, _ := regexp.MatchString(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`, str)
	return isMatch
}

func ValdateCity(str string) bool {
	_, cityExist := city[str]
	return cityExist
}

// func ValidateUsername(str string) bool {
// 	isMatch, _ := regexp.MatchString(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`, str)
// 	return isMatch
// }

// func ValidateCity(str string) bool {
// 	isMatch, _ := regexp.MatchString(`^([А-Я]{1,100}){1,5}$`, str)
// 	return isMatch
// }

// func ValidatePassword(str string) bool {
// 	isMatch, _ := regexp.MatchString(`/(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z!@#$%^&*]{8,}/g`, str)
// 	return isMatch
// }
