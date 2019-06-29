package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

var stopServerChannel = make(chan bool, 1)

func startServer() {
	if handleFunction == nil {
		fmt.Println("ERROR: Cannot start - no handle function!")
		return
	}

	fmt.Println("Server STARTING")

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", handleFunction)
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: httpMux,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		},
	}
	go func() {
		if cfg.Server.TLSEnabled {
			fmt.Println("    ...with TLS")
			if err := httpServer.ListenAndServeTLS(cfg.Server.TLSFileCrt, cfg.Server.TLSFileKey); err != nil && err != http.ErrServerClosed {
				fmt.Printf("ERROR: %T, %q", err, err)
			}
		} else {
			fmt.Println("    ...without TLS")
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("ERROR: %T, %q", err, err)
			}
		}
		fmt.Println("Server STOPPED")
	}()

	go func() {
		<-stopServerChannel
		fmt.Println("Server STOPPING")
		httpServer.Shutdown(context.Background())
	}()

	fmt.Println("Server STARTED")
}

func stopServer() {
	fmt.Println("...we will try to stop the server")
	stopServerChannel <- true
	time.Sleep(500000000)
}
