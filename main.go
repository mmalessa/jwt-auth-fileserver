package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mmalessa/go-http-fileserver-jwt/config"
	"github.com/mmalessa/go-http-fileserver-jwt/httpserver"
)

func main() {
	fmt.Println("start")

	myPid := os.Getpid()
	fmt.Printf("My PID: %d\n", myPid)

	config := config.Init()
	fmt.Println(config)

	httpServer := httpserver.Init()
	httpServer.Port = config.ServerPort
	httpServer.HandleFunction = httpHandle

	fmt.Println(httpServer)

	httpServer.Start()

	stopServer := make(chan os.Signal, 1)
	signal.Notify(stopServer,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	sig := <-stopServer
	fmt.Printf("%T %s\n", sig, sig)
	httpServer.Stop()

}

func httpHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("akuku")
}
