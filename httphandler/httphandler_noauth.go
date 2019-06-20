package httphandler

import (
	"fmt"
	"net/http"
)

func (c *config) HttpHandleNoAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: HttpHandleNoAuth")
}
