package httpErrDescr

const (
	AUTH_REQUIRED      = "для этого действия необходимо авторизоваться"
	BAD_AUTH_TOKEN     = "невалидный токен авторизации"
	NO_SUCH_RESTAURANT = "не существует такого ресторана"
	NO_SUCH_CODE_INFO  = "не найдены данные для проверки кода пользователя"
	BAD_REQUEST_BODY   = "не правильное тело запроса"
	SERVER_ERROR       = "ошибка на сервере"
	WRONG_AUTH_CODE    = "неверный код для входа"
	INVALID_DATA       = "переданы невалидные данные"
	ALREADY_AUTHORIZED = "пользователь уже авторизован"
)
