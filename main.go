package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mmalessa/go-http-fileserver-jwt/config"
	"github.com/mmalessa/go-http-fileserver-jwt/httphandler"
	"github.com/mmalessa/go-http-fileserver-jwt/httpserver"
)

func main() {
	fmt.Println("STARTING")

	myPid := os.Getpid()
	fmt.Printf("My PID is: %d\n", myPid)

	config := config.Init()
	fmt.Print("config: ")
	fmt.Println(config)

	httpServer := httpserver.Init()
	httpServer.Port = config.ServerPort
	httpServer.HandleFunction = httphandler.HttpHandle
	httpServer.Tls = config.ServerTls
	httpServer.FileCrt = config.ServerFileCrt
	httpServer.FileKey = config.ServerFileKey

	fmt.Print("http server: ")
	fmt.Println(httpServer)

	httpServer.Start()

	stopServer := make(chan os.Signal, 1)
	signal.Notify(stopServer, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-stopServer
	fmt.Printf("Ask for stop with signal: %T %s\n", sig, sig)
	httpServer.Stop()
}
