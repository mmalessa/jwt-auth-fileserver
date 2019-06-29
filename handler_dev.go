package main

import (
	"fmt"
	"net/http"
)

func handleDev(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "TEST PAGE\nhttps://github.com/mmalessa/jwt-auth-fileserver\n\n\n")
	fmt.Fprintf(w, "REQUEST\n %+v\n\n", *r)
	fmt.Fprintf(w, "CONFIG\n %+v\n\n", *cfg)
}
