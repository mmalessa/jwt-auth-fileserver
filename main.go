package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type handleFunctionType func(w http.ResponseWriter, r *http.Request)

var cfg *Config
var err error
var handleFunction handleFunctionType

func main() {
	fmt.Println("STARTING")

	myPid := os.Getpid()
	fmt.Printf("My PID is: %d\n", myPid)

	configFile := "config.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	cfg, err = loadConfig(configFile)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: %v", err))
		return
	}
	fmt.Printf("%+v\n", *cfg)

	handleFunction = getHandleFunction(cfg.Handler.AuthType)

	_ = handleFunction

	fmt.Print("http server: ")

	startServer()

	stopServerChannel := make(chan os.Signal, 1)
	signal.Notify(stopServerChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-stopServerChannel
	fmt.Printf("Ask for stop with signal: %T %s\n", sig, sig)
	stopServer()
}
