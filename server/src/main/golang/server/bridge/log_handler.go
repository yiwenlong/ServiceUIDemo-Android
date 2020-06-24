package bridge

type ILogHandler interface {
	onLog(message string)
}
