package configRouting

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type ServerHandlers struct {
	// UserHandler *userHttp.UserHandler
	/// ...
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	// v1Prefix := "api/v1/"
	router.POST("/", hello)
	// router.POST(v1Prefix+"login", sh.UserHandler.Login)
	// router.POST(v1Prefix+"logout", sh.UserHandler.Logout)
	// router.POST(v1Prefix+"register", sh.UserHandler.Register)
	////...
}

// TODO: delete this
func hello(ctx echo.Context) error {
	fmt.Println(ctx)
	fmt.Println("---------")
	fmt.Println(ctx.Request())
	fmt.Println("---------")

	ctx.JSON(200, "hi")
	fmt.Println(ctx.Response())

	return nil
}
