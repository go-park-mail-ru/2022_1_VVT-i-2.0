package notification

type Notificator interface {
	SendCode(phone string, code string) error
}
