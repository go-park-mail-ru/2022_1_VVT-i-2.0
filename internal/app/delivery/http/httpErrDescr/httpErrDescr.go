package httpErrDescr

import "github.com/labstack/echo/v4"

const (
	AUTH_REQUIRED                  = "для этого действия необходимо авторизоваться"
	BAD_AUTH_TOKEN                 = "невалидный токен авторизации"
	NO_SUCH_RESTAURANT             = "не существует такого ресторана"
	NO_SUCH_DISHES                 = "не существует подходящих блюд"
	NO_SUCH_RESTAURANTS            = "не существует подходящих ресторанов"
	NO_SUCH_USER                   = "не существует такого пользователя"
	THIS_ORDER_DOESNOT_BELONG_USER = "этот заказ не принадлежит текущему пользователю"
	BAD_ORDER_ID                   = "неалидный  номер заказа"
	NO_SUCH_CODE_INFO              = "не найдены данные для проверки кода пользователя"
	SUCH_USER_ALREADY_EXISTS       = "пользователь с таким телефоном или почтой уже зарегистрирован"
	BAD_REQUEST_BODY               = "неправильное тело запроса"
	BAD_IMAGE                      = "ошибка при открытии загруженной картинки"
	SERVER_ERROR                   = "ошибка на сервере"
	CREATING_ORDER                 = "такой заказ не может быть создан"
	WRONG_AUTH_CODE                = "неверный код для входа"
	INVALID_DATA                   = "переданы невалидные данные"
	ALREADY_AUTHORIZED             = "пользователь уже авторизован"
	INVALID_CSRF                   = "невалидный csrf-токен"
)

func NewHTTPError(ctx echo.Context, httpStatusCode int, descr string) error {
	ctx.Response().Status = httpStatusCode
	return echo.NewHTTPError(httpStatusCode, descr)
}
