package main

func getHandleFunction(authType string) handleFunctionType {
	switch authType {
	case "noauth":
		return handleNoAuth
	case "jwt":
		return handleJwt
	default:
		return handleDev
	}
}
