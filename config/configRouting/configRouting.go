package configRouting

import (
	"github.com/labstack/echo/v4"
)

type ServerHandlers struct {
	// UserHandler *userHttp.UserHandler
	/// ...
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	// v1Prefix := "api/v1/"
	// router.POST(v1Prefix+"login", sh.UserHandler.Login)
	// router.POST(v1Prefix+"logout", sh.UserHandler.Logout)
	// router.POST(v1Prefix+"register", sh.UserHandler.Register)
	////...
}
