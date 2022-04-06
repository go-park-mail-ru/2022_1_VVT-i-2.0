package errorDescription

// // ошибки, которые можно показать пользователю [1,+inf]
// const (
// 	AUTH_REQUIRED int = iota + 1
// 	NO_SUCH_RESTAURANT
// )

// // ошибки, которые не стоит в этом виде показывать пользователю [-inf;0]
// const (
// 	BAD_REQUEST_BODY int = -iota
// 	SERVER_ERROR     int = -iota
// )

const (
	AUTH_REQUIRED_DESCR      = "для этого действия необходимо войти авторизоваться"
	BAD_AUTH_TOKEN           = "не валидный токен авторизации"
	NO_SUCH_RESTAURANT_DESCR = "не существует такого ресторана"
	BAD_REQUEST_BODY_DESCR   = "не правильное тело запроса"
	SERVER_ERROR_DESCR       = "ошибка на сервера"
)

// var codeDescr = map[int]string{
// 	AUTH_REQUIRED:      AUTH_REQUIRED_DESCR,
// 	NO_SUCH_RESTAURANT: NO_SUCH_RESTAURANT_DESCR,
// 	BAD_REQUEST_BODY:   BAD_REQUEST_BODY_DESCR,
// 	SERVER_ERROR:       SERVER_ERROR_DESCR,
// }

// type Error struct {
// 	Description string
// 	Code        int
// }

// func (e Error) Error() string {
// 	return fmt.Sprintf("error with code %d description: %s", e.Code, e.Description)
// }

// func NewError(eCode int) *Error {
// 	descr := codeDescr[eCode]
// 	if descr == "" {
// 		return nil
// 	}
// 	return &Error{
// 		Code:        eCode,
// 		Description: descr,
// 	}
// }
