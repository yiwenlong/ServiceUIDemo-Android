package bridge

import "github.com/yiwenlong/server"

func BootServer(host string, port int) {
	server.BootServer(host, port)
}

func RegisterLogHandler(logHandler *ILogHandler) {

}

func StopServer(port int) {
	server.StopServer(port)
}

func StopAllServer() {
	server.StopAllServer()
}