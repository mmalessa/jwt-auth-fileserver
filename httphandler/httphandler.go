package httphandler

import (
	"fmt"
	"html"
	"net/http"
)

type config struct {
	AuthApiEndpoint    string
	AuthApiType        string
	FilesRootDirectory string
}

func Init() *config {
	c := new(config)
	return c
}

func (c *config) HttpHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello: %q \n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "Header: %s\n", r.Header)
	fmt.Fprintf(w, "\n\nRequest: %s", r)
	fmt.Fprintf(w, "\n\nConfig: %s", c)
	fmt.Println(r.URL.Path)
}
