package main

import (
	"fmt"
	"github.com/yiwenlong/server/bridge"
	"time"
)

type ServerHandler struct {
}

func (lh *ServerHandler) Log(message string) {
	fmt.Printf("MyLog ==>  %s", message)
}

func (lh *ServerHandler) OnServerStart() {
	fmt.Printf("MyLog ==> OnServerStart")
}

func (lh *ServerHandler) OnServerStop() {
	fmt.Printf("MyLog ==> OnServerStop")
}

func (lh *ServerHandler) OnServerError(msg string) {
	fmt.Printf("MyLog ==> OnServerStop")
}

func main() {
	bridge.RegisterLogHandler(&ServerHandler{})
	bridge.InitServer("localhost", 8080)
	bridge.RegisterServerListener(&ServerHandler{})
	bridge.BootServer()
	time.Sleep(1 * time.Minute)
	bridge.StopServer()
}
