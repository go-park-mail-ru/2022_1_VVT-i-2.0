package httpErrDescr

import "github.com/labstack/echo/v4"

const (
	AUTH_REQUIRED                  = "Вам необходимо авторизоваться"
	BAD_AUTH_TOKEN                 = "Невалидный токен авторизации"
	NO_SUCH_RESTAURANT             = "Извините, такого ресторана не существует"
	ORDER_ALREADY_REVIEWED         = "Извините, но Вы уже оставляли отзыв на этот заказ"
	NO_SUCH_DISHES                 = "Не существует подходящих блюд"
	NO_SUCH_RESTAURANTS            = "Не существует подходящих ресторанов"
	NO_SUCH_USER                   = "Этот пользователь не зарегистрирован"
	THIS_ORDER_DOESNOT_BELONG_USER = "Этот заказ не принадлежит Вам"
	BAD_ORDER_ID                   = "Вы передали неправильный  номер заказа"
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
	INVALID_ADDRESS                = "Cлишком длинный адрес, либо он содержит неожидаемые символы."
	EMPTY_CART                     = "Извините, Ваша корзина пуста."
	INVALID_COMMENT                = "Извините, Ваш комментарий слишком длинный, оставьте 1000 символов."
	INVALID_SLUG                   = "Извините, данного ресторана не существует."
	INVALID_PROMOCODE              = "Невалидный промокод, он может содержать только латинские буквы и цифры."
	INVALID_ENTRANCE               = "Вы указали некорректный подъезд."
	INVALID_INTERCOM               = "Вы указали некорректный код домофона."
	INVALID_FLOOR                  = "Вы указали некорректный этаж."
	INVALID_FLAT                   = "Вы указали некорректный номер квартиры."
	INVALID_CATEGORY               = "Невалидная категория ресторанов, категория может содержать от 2 до 25 цифр и букв."
	INVALID_QUERY                  = "Извините, но название ресторана не может превышать 100 символов."
	INVALID_REVIEW                 = "Пожалуйста, введите от 1 до 1000 символов."
	INVALID_RATING                 = "Оценка должна быть целым числом от 1 до 5."
	INVALID_PHONE                  = "Извините, мы обслуживаем только российские номера."
	INVALID_AUTHCODE               = "Вы указали неправильный код. Код должен содержать 4 цифры."
	INVALID_EMAIL                  = "Вы указали некорректный email."
	INVALID_NAME                   = "Вы указали некорректное имя. Имя может содержать от 2 до 25 цифр и букв."
)

func NewHTTPError(ctx echo.Context, httpStatusCode int, descr string) error {
	ctx.Response().Status = httpStatusCode
	return echo.NewHTTPError(httpStatusCode, descr)
}
