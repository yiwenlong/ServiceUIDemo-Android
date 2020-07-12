package bridge

import (
	"github.com/yiwenlong/server"
	"log"
	"os"
)

var myServer *server.MyServer

type LogHandlerWrapper struct {
	logHandler ILogHandler
}

func (lh *LogHandlerWrapper) Write(log []byte) (int, error) {
	if lh.logHandler != nil {
		lh.logHandler.Log(string(log))
	}
	return len(log), nil
}

func BootServer() {
	if myServer == nil {
		log.Fatalf("Server not init!")
	}
	myServer.Boot()
}

func InitServer(host string, port int) {
	if myServer != nil {
		myServer.Stop()
	}
	myServer = server.NewMyServer(host, port)
}

func RegisterLogHandler(logHandler ILogHandler) {
	log.SetOutput(&LogHandlerWrapper{
		logHandler: logHandler,
	})
}

func UnRegisterLogHandler() {
	log.SetOutput(os.Stderr)
}

func RegisterServerListener(listener MyServerListener) {
	if myServer == nil {
		log.Fatalf("Server not init!")
	}
	myServer.RegisterListener(listener)
}

func ClearServerListeners() {
	if myServer != nil {
		myServer.ClearListeners()
	}
}

func StopServer() {
	if myServer == nil {
		log.Fatalf("Server not init!")
	}
	myServer.Stop()
	myServer = nil
}

func ServerIsRuning() bool {
	if myServer == nil {
		return false
	}
	return myServer.IsRunning()
}
