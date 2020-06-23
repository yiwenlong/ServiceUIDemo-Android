package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

var servers map[int]*http.Server

func BootServer(host string, port int) {

	if servers == nil {
		servers = make(map[int]*http.Server, 10)
	}

	address := fmt.Sprintf("%s:%d", host, port)
	srv := http.Server{Addr: address}

	log.Printf("Server start listening on %s\n", address)

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	go func() {
		defer func() {
			delete(servers, port)
			log.Printf("Server finish.")
		}()
		servers[port] = &srv
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
}

func IsServerRunning(port int) bool {
	_, ok := servers[port]
	return ok
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		log.Fatalf("Send message to client failed: %v\n", err)
	}
	log.Printf("URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, err := fmt.Fprintf(w, "Count %d\n", count)
	if err != nil {
		log.Fatalf("Send message to client failed: %v\n", err)
	}
	mu.Unlock()
	log.Printf("Count %d\n", count)
}


func StopServer(port int) error {
	if !IsServerRunning(port) {
		return nil
	}
	server := servers[port]
	log.Printf("Start shut down server on port: %d\n", port)
	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Server shut down failed with error: %v\n", err)
		return err
	}
	delete(servers, port)
	log.Printf("Server shut down.")
	return nil
}

func StopAllServer() error {
	for port, _ := range servers {
		if err := StopServer(port); err != nil {
			return err
		}
	}
	return nil
}