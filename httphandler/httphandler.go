package httphandler

import (
	"net/http"
)

type handleFn func(w http.ResponseWriter, r *http.Request)

type config struct {
	authApiEndpoint    string
	authApiType        string
	filesRootDirectory string
	HandleFunction     handleFn
}

func Init(filesRootDirectory string, authApiType string, authApiEndpoint string) *config {
	c := new(config)
	c.authApiEndpoint = authApiEndpoint
	c.authApiType = authApiType
	c.filesRootDirectory = filesRootDirectory
	c.HandleFunction = c.getHandleFunction()
	return c
}

func (c *config) getHandleFunction() handleFn {
	switch c.authApiType {
	case "noauth":
		return c.HttpHandleNoAuth
	case "jwt":
		return c.HttpHandleJwt
	default:
		return c.HttpHandleDev
	}
}
