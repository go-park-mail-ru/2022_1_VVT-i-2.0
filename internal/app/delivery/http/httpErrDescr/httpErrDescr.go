package httpErrDescr

const (
	AUTH_REQUIRED      = "для этого действия необходимо авторизоваться"
	BAD_AUTH_TOKEN     = "невалидный токен авторизации"
	NO_SUCH_RESTAURANT = "не существует такого ресторана"
	NO_SUCH_AUTH_DATA  = "не найдены данные для авторизации пользователя"
	BAD_REQUEST_BODY   = "не правильное тело запроса"
	SERVER_ERROR       = "ошибка на сервера"
	WRONG_AUTH_CODE    = "неверный код для входа"

	// невалидные данные от пользователя
	INVALID_DATA = "переданы невалидные данные"
)
