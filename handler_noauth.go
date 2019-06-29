package main

import (
	"fmt"
	"net/http"
)

func handleNoAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: HttpHandleNoAuth")
}
