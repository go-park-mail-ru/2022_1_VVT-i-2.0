package configRouting

import (
	restaurantsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/delivery/http"
	userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery/http"
	"github.com/labstack/echo/v4"
)

type ServerHandlers struct {
	UserHandler *userHandler.UserHandler
	RestaurantsHandler *restaurantsHandler.RestaurantsHandler
	/// ...
}

const (
	v1Prefix = "/api/v1/"
)

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	router.POST(v1Prefix+"login", sh.UserHandler.Login)
	router.GET(v1Prefix+"logout", sh.UserHandler.Logout)
	router.POST(v1Prefix+"register", sh.UserHandler.Register)
	router.POST(v1Prefix+"update", sh.UserHandler.UpdateUser)
	router.POST(v1Prefix+"sendcode", sh.UserHandler.SendCode)
	router.GET(v1Prefix+"user", sh.UserHandler.GetUser)
	router.GET(v1Prefix+"restaurants", sh.RestaurantsHandler.GetAllRestaurants)
	////...
}
