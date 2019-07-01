package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type handleFunctionType func(w http.ResponseWriter, r *http.Request)

type ErrorMessage struct {
	Code    string
	Message string
}

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

	cfgM, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Print("CONFIG: ")
	fmt.Println(string(cfgM))

	handleFunction = getHandleFunction(cfg.Handler.AuthType)

	startServer()

	stopServerChannel := make(chan os.Signal, 1)
	signal.Notify(stopServerChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-stopServerChannel
	fmt.Printf("Ask for stop with signal: %T %s\n", sig, sig)
	stopServer()
}
