package main

import (
	"fmt"
	"github.com/yiwenlong/server/bridge"
	"time"
)

type LogHandler struct {
}

func (lh *LogHandler) Log(message string) {
	fmt.Printf("MyLog ==>  %s", message)
}

func main() {
	bridge.RegisterLogHandler(&LogHandler{})
	bridge.BootServer("0.0.0.0", 8080)
	time.Sleep(1 * time.Minute)
	bridge.StopServer(8080)
}
