package main

import (
	"github.com/yiwenlong/server"
	"log"
	"time"
)

func main()  {
	server.BootServer("localhost", 8080)
	time.Sleep(5 * time.Second)
	log.Printf("Server is running: %v\n", server.IsServerRunning(8080))
	time.Sleep(5 * time.Minute)
	log.Printf("Server is running: %v\n", server.IsServerRunning(8080))
	server.StopServer(8080)
}
