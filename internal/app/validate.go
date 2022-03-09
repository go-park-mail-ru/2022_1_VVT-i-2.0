package serv

import (
	// "github.com/gorilla/mux"
	// "go.uber.org/zap"

	"fmt"

	"github.com/dongri/phonenumber"
)

var number string = "+358401231234"
var defaultRegion string = "RU"

func validatePhone(str string) bool {
	possibleNumber := phonenumber.ParseWithLandLine(number, defaultRegion)
	fmt.Printf("The number %s is possible phone number in region %s: %v", number, defaultRegion, possibleNumber)
	return (possibleNumber != "")
}
