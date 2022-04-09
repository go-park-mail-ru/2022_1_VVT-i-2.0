package notification

type Notificator interface {
	// ...
	// Send(phone string, msg string, encoding string) error
	SendCode(phone string, code string) error
}
