package serv

import (
	"github.com/dongri/phonenumber"
)

var number string = "+358401231234"
var defaultRegion string = "RU"

func validatePhone(str string) bool {
	possibleNumber := phonenumber.ParseWithLandLine(number, defaultRegion)
	return (possibleNumber != "")
}
