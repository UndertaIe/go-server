package logger

type Log interface {
	Debug(...any)
	Info(...any)
	Warn(...any)
	Error(...any)
}
