package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var stopMe = make(chan bool, 1)

type config struct {
	Port                int
	Tls                 bool
	Timeout             int
	DialTimeout         int
	TlsHandshakeTimeout int
	HandleFunction      func(w http.ResponseWriter, r *http.Request)
	FileCrt             string
	FileKey             string
}

func Init() *config {
	c := new(config)
	return c
}

func (c *config) Start() {
	if c.HandleFunction == nil {
		fmt.Println("ERROR: Cannot start - no handle function!")
		return
	}
	fmt.Println("Server STARTING")

	httpMux := http.NewServeMux()
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: httpMux,
	}
	httpMux.HandleFunc("/", c.HandleFunction)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ERROR: %T, %q", err, err)
		}
		fmt.Println("Server STOPPED")
	}()

	go func() {
		<-stopMe
		fmt.Println("Server STOPPING")
		httpServer.Shutdown(context.Background())
	}()

	fmt.Println("Server STARTED")
}

func (c *config) Stop() {
	fmt.Println("...we will try to stop the server")
	stopMe <- true
	time.Sleep(500000000)
}

// FIXME
func startTLS() {
	// http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
