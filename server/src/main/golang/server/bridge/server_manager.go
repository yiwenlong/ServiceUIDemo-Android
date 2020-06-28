package bridge

import (
	"github.com/yiwenlong/server"
	"log"
)

type LogHandlerWrapper struct {
	logHandler ILogHandler
}

func (lh *LogHandlerWrapper) Write(log []byte) (int, error) {
	if lh.logHandler != nil {
		lh.logHandler.Log(string(log))
	}
	return len(log), nil
}

func BootServer(host string, port int) {
	server.BootServer(host, port)
}

func RegisterLogHandler(logHandler ILogHandler) {
	log.SetOutput(&LogHandlerWrapper{
		logHandler: logHandler,
	})
}

func StopServer(port int) {
	server.StopServer(port)
}

func StopAllServer() {
	server.StopAllServer()
}
