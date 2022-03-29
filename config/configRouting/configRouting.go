package configRouting

import (
	"fmt"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/labstack/echo/v4"
)

type ServerHandlers struct {
	// UserHandler *userHttp.UserHandler
	/// ...
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	// v1Prefix := "api/v1/"
	router.GET("/", hello)
	// TODO: set auth-mw and auth-opt-mw to urls
	// router.POST(v1Prefix+"login", sh.UserHandler.Login)
	// router.POST(v1Prefix+"logout", sh.UserHandler.Logout)
	// router.POST(v1Prefix+"register", sh.UserHandler.Register)
	////...
}

// TODO: delete this
func hello(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	fmt.Println("--user----")
	fmt.Println("user from ctx", user)
	logger := middleware.GetLoggerFromCtx(ctx)
	fmt.Println("logger from ctx", logger)
	fmt.Println("reqId from ctx", middleware.GetRequestIdFromCtx(ctx))
	fmt.Println(ctx)
	fmt.Println("---------")
	fmt.Println(ctx.Request())
	fmt.Println("---------")

	(*logger).Infow("I am in hello func")
	if user != nil {
		ctx.JSON(200, struct {
			Msg    string
			YourId models.UserId
		}{Msg: "hi", YourId: user.Id})
		return nil
	}
	ctx.JSON(200, struct {
		Msg string `json:"msg"`
	}{Msg: "hi incognito!"})
	fmt.Println(ctx.Response())

	return nil
}
