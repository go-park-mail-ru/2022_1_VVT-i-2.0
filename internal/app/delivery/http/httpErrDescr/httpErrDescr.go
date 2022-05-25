package httpErrDescr

import "github.com/labstack/echo/v4"

const (
	AUTH_REQUIRED                  = "Вам необходимо авторизоваться"
	BAD_AUTH_TOKEN                 = "Невалидный токен авторизации"
	NO_SUCH_RESTAURANT             = "Не существует такого ресторана"
	NO_SUCH_DISHES                 = "Не существует подходящих блюд"
	NO_SUCH_RESTAURANTS            = "Не существует подходящих ресторанов"
	NO_SUCH_USER                   = "Этот пользователь не зарегистрирован"
	THIS_ORDER_DOESNOT_BELONG_USER = "Этот заказ не принадлежит Вам"
	BAD_ORDER_ID                   = "Неправильный  номер заказа"
	NO_SUCH_CODE_INFO              = "Не найдены данные для проверки кода пользователя"
	SUCH_USER_ALREADY_EXISTS       = "Пользователь с таким телефоном или почтой уже зарегистрирован"
	BAD_REQUEST_BODY               = "Неправильное тело запроса"
	BAD_IMAGE                      = "Ошибка при открытии загруженной картинки"
	SERVER_ERROR                   = "Ошибка на сервере"
	CREATING_ORDER                 = "Такой заказ не может быть создан"
	NO_SUCH_ADDRESS                = "Неправильно указан адрес"
	WRONG_AUTH_CODE                = "Неверный код для входа"
	INVALID_DATA                   = "Переданы невалидные данные"
	ALREADY_AUTHORIZED             = "Пользователь уже авторизован"
	INVALID_CSRF                   = "Невалидный csrf-токен"
)

func NewHTTPError(ctx echo.Context, httpStatusCode int, descr string) error {
	ctx.Response().Status = httpStatusCode
	return echo.NewHTTPError(httpStatusCode, descr)
}
