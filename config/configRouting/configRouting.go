package configRouting

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery/http"
	"github.com/labstack/echo/v4"
)

type ServerHandlers struct {
	UserHandler *userHandler.UserHandler
	/// ...
}

const (
	V1Prefix = "/api/v1/"
)

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	// TODO: удалить
	router.GET("/", noErr)
	router.GET("/err", err)
	router.GET("/panic", panicH)
	router.HTTPErrorHandler = router.DefaultHTTPErrorHandler
	// router.GET("/h", hello)

	router.POST(V1Prefix+"login", sh.UserHandler.Login)
	// router.POST(v1Prefix+"logout", sh.UserHandler.Logout)
	// router.POST(v1Prefix+"register", sh.UserHandler.Register)
	////...
}

// TODO: delete this
func err(ctx echo.Context) error {
	fmt.Println("=====in err func=====")
	return echo.NewHTTPError(http.StatusBadRequest, servErrors.BAD_REQUEST_BODY_DESCR)
}

func noErr(ctx echo.Context) error {
	fmt.Println("=====in no-err func=====")
	return ctx.JSON(200, struct {
		Data string `json:"data-key"`
	}{Data: "data-val"})
}

func panicH(ctx echo.Context) error {
	fmt.Println("=====in panic func=====")
	panic("I panic")
	// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
}

/*
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

	// err := echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	// fmt.Printf(err.Error())
	return nil

	// return echo.NewHTTPError(http.StatusUnauthorized, errorDescription.AUTH_REQUIRED_DESCR)
}
*/
