package httphandler

import (
	"fmt"
	"html"
	"net/http"
)

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello: %q \n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "Header: %s\n", r.Header)
	fmt.Fprintf(w, "\n\n%s", r)
	fmt.Println(r.URL.Path)
}
