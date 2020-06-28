package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type StartListener interface {
	OnServerStart()
}

type StopListener interface {
	OnServerStop()
}

type ErrorListener interface {
	OnServerError(msg string)
}

type MyServer struct {
	address        string
	port           int
	mutex          sync.Mutex
	count          int
	server         *http.Server
	startListeners []StartListener
	stopListeners  []StopListener
	errorListeners []ErrorListener
	isRunning      bool
}

func NewMyServer(addr string, port int) *MyServer {
	return &MyServer{
		address: addr,
		port:    port,
		mutex:   sync.Mutex{},
		count:   0,
	}
}

func (serv *MyServer) Boot() {
	address := fmt.Sprintf("%s:%d", serv.address, serv.port)
	if serv.isRunning {
		log.Fatalf("Server already started: %s", address)
		return
	}
	serv.server = &http.Server{Addr: address}
	http.HandleFunc("/", serv.handler)
	http.HandleFunc("/count", serv.counter)
	log.Printf("Start boot server on %s\n", address)

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		for sig := range signalChan {
			log.Printf("Received signal: %d (%s)", sig, sig)
			serv.Stop()
		}
	}()

	go func() {
		defer func() {
			serv.isRunning = false
			log.Printf("Server finished.")
			for i := 0; i < len(serv.stopListeners); i++ {
				serv.stopListeners[i].OnServerStop()
			}
		}()
		if err := serv.server.ListenAndServe(); err != http.ErrServerClosed {
			for i := 0; i < len(serv.errorListeners); i++ {
				serv.errorListeners[i].OnServerError(err.Error())
			}
			log.Fatalf("Server failed: %v", err)
		}
	}()
	serv.isRunning = true
	log.Printf("Server listening on %s\n", address)
	for i := 0; i < len(serv.startListeners); i++ {
		serv.startListeners[i].OnServerStart()
	}
}

func (serv *MyServer) Stop() {
	if serv.server == nil || !serv.isRunning {
		log.Fatalf("Server is NOT running")
	}
	if err := serv.server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Server shut down failed with error: %v\n", err)
	}
	log.Printf("Server shut down.")
}

func (serv *MyServer) handler(w http.ResponseWriter, r *http.Request) {
	serv.mutex.Lock()
	serv.count++
	serv.mutex.Unlock()
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		log.Fatalf("Send message to client failed: %v\n", err)
		return
	}
	log.Printf("URL.Path = %q\n", r.URL.Path)
}

func (serv *MyServer) counter(w http.ResponseWriter, r *http.Request) {
	serv.mutex.Lock()
	_, err := fmt.Fprintf(w, "Count %d\n", serv.count)
	if err != nil {
		log.Fatalf("Send message to client failed: %v\n", err)
	}
	serv.mutex.Unlock()
	log.Printf("Count %d\n", serv.count)
}

func (serv *MyServer) AddStartListener(listener StartListener) {
	if serv.startListeners == nil {
		serv.startListeners = []StartListener{}
	}
	serv.startListeners = append(serv.startListeners, listener)
}

func (serv *MyServer) AddStopListener(listener StopListener) {
	if serv.stopListeners == nil {
		serv.stopListeners = []StopListener{}
	}
	serv.stopListeners = append(serv.stopListeners, listener)
}

func (serv *MyServer) AddErrorListener(listener ErrorListener) {
	if serv.errorListeners == nil {
		serv.errorListeners = []ErrorListener{}
	}
	serv.errorListeners = append(serv.errorListeners, listener)
}

func (serv *MyServer) IsRunning() bool {
	return serv.isRunning
}
