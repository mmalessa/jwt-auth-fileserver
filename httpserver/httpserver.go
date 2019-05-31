package httpserver

import (
	"context"
	"crypto/tls"
	"fmt"
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
	httpMux.HandleFunc("/", c.HandleFunction)
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
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
		if c.Tls {
			fmt.Println("    ...with TLS")
			if err := httpServer.ListenAndServeTLS(c.FileCrt, c.FileKey); err != nil && err != http.ErrServerClosed {
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
