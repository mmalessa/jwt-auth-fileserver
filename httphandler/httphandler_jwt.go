package httphandler

import (
	"fmt"
	"net/http"
)

func (c *config) HttpHandleJwt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: HttpHandleJwt")
}
