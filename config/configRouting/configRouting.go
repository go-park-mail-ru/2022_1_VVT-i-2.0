package configRouting

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	_ "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/docs"
	suggestHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/delivery/http"
	commentHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments/delivery/http"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	dishesHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/delivery/http"
	orderHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/delivery/http"
	promocodeHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode/delivery/http"
	recommendationsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/delivery/http"
	restaurantsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/delivery/http"
	userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery/http"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerHandlers struct {
	UserHandler            *userHandler.UserHandler
	RestaurantsHandler     *restaurantsHandler.RestaurantsHandler
	DishesHandler          *dishesHandler.DishesHandler
	SuggsHandler           *suggestHandler.SuggsHandler
	OrderHandler           *orderHandler.OrderHandler
	CommentsHandler        *commentHandler.CommentsHandler
	RecommendstionsHandler *recommendationsHandler.RecommendationsHandler
	PromocodeHandler       *promocodeHandler.PromocodesHandler
}

const (
	v1Prefix = "/api/v1/"
)

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo, mw *middleware.CommonMiddleware, corsCfg *config.CorsConfig, csrfCfg *config.CsrfConfig) {
	router.HTTPErrorHandler = mw.ErrorHandler
	router.Use(mw.PanicMiddleware)
	router.Use(echoMiddleware.CORSWithConfig(getCorsConfig(corsCfg)))
	mwChain := []echo.MiddlewareFunc{
		mw.RequestIdMiddleware,
		mw.AccessLogMiddleware,
		mw.AuthMiddleware,
		middleware.CSRFWithConfig(getCsrfConfig(csrfCfg)),
		mw.PanicMiddleware,
		middleware.CsrfSetHeader,
	}

	router.GET("metrics", echo.WrapHandler(promhttp.Handler()))
	router.Static("/static", "static")
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.POST(v1Prefix+"login", sh.UserHandler.Login, mwChain...)
	router.GET(v1Prefix+"logout", sh.UserHandler.Logout, mwChain...)
	router.POST(v1Prefix+"register", sh.UserHandler.Register, mwChain...)
	router.POST(v1Prefix+"update", sh.UserHandler.UpdateUser, mwChain...)
	router.POST(v1Prefix+"send_code", sh.UserHandler.SendCode, mwChain...)
	router.GET(v1Prefix+"user", sh.UserHandler.GetUser, mwChain...)
	router.GET(v1Prefix+"suggest", sh.SuggsHandler.Suggest, mwChain...)
	router.POST(v1Prefix+"order", sh.OrderHandler.CreateOrder, mwChain...)
	router.GET(v1Prefix+"orders", sh.OrderHandler.GetUserOrders, mwChain...)
	router.GET(v1Prefix+"order_statuses", sh.OrderHandler.GetUserOrderStatuses, mwChain...)
	router.GET(v1Prefix+"order/:orderId", sh.OrderHandler.GetUserOrder, mwChain...)
	router.GET(v1Prefix+"comments/:slug", sh.CommentsHandler.GetRestaurantComments, mwChain...)
	router.POST(v1Prefix+"comment", sh.CommentsHandler.AddRestaurantComment, mwChain...)
	router.GET(v1Prefix+"restaurants", sh.RestaurantsHandler.GetAllRestaurantsMain, mwChain...)
	router.GET(v1Prefix+"promo", sh.PromocodeHandler.GetAllPromocodes, mwChain...)
	router.GET(v1Prefix+"", sh.RestaurantsHandler.GetAllRestaurants, mwChain...)
	router.GET(v1Prefix+"restaurant/:slug", sh.DishesHandler.GetDishesByRestaurants, mwChain...)
	router.POST(v1Prefix+"recommendations", sh.RecommendstionsHandler.GetRecommendations, mwChain...)
}
