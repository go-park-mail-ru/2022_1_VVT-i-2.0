package recommendationsHandler

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
)

type RestaurantsHandler struct {
	Ucase         recommendations.Ucase
	StaticManager staticManager.FileManager
}

func NewRestaurantsHandler(ucase restaurants.Ucase, staticManager staticManager.FileManager) *RestaurantsHandler {
	return &RestaurantsHandler{
		Ucase:         ucase,
		StaticManager: staticManager,
	}
}