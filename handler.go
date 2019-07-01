package main

func getHandleFunction(authType string) handleFunctionType {
	switch authType {
	case "none":
		return handleNoAuth
	case "jwt":
		return handleJwt
	default:
		return handleDev
	}
}
