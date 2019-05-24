package httpserver

import (
	"fmt"
	"net/http"
	"time"
)

var stopServerChannel = make(chan bool)

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
	fmt.Println("HTTP server start")

	go func() {
		<-stopServerChannel
		fmt.Println("HTTP server stop")
	}()
}

func (c *config) Stop() {
	fmt.Println("...try stop")
	stopServerChannel <- true
	time.Sleep(1000000000)
}

// func main() {
//     http.HandleFunc("/hello", HelloServer)
//     err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
//     if err != nil {
//         log.Fatal("ListenAndServe: ", err)
//     }
// }
