package main

import (
	"fmt"
	"github.com/yiwenlong/server"
	"log"
	"time"
)

type Listener struct {
}

func (l *Listener) OnServerStart() {
	fmt.Printf(">>>> start <<<<\n")
}

func (l *Listener) OnServerStop() {
	fmt.Printf(">>>> stop <<<<\n")
}

func (l *Listener) OnServerError() {
	fmt.Printf(">>>> error <<<<\n")
}

func main() {
	myServer := server.NewMyServer("localhost", 8080)
	listener := Listener{}
	myServer.AddStartListener(&listener)
	myServer.AddStopListener(&listener)
	myServer.Boot()
	time.Sleep(time.Second)
	log.Printf("Server is running: %v\n", myServer.IsRunning())
	time.Sleep(5 * time.Second)
	log.Printf("Server is running: %v\n", myServer.IsRunning())
	myServer.Stop()
	time.Sleep(time.Second)
}
