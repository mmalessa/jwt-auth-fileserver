package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func handleJwt(w http.ResponseWriter, r *http.Request) {
	urlPath := html.EscapeString(r.URL.Path)
	filename := strings.TrimPrefix(urlPath, "/")
	filePath := string(cfg.Handler.RootDirectory + "/" + filename)
	_, err := os.Stat(filePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "404", Message: http.StatusText(404)})
		log.Println("NOT FOUND (" + urlPath + ")")
		return
	}

	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
		log.Println("NO TOKEN FOUND (" + urlPath + ")")
		return
	}
	request, err := http.NewRequest("GET", cfg.Handler.JwtTestEndpoint+"?file="+filename, nil)
	// FIXME - add token structure validation ...or something
	tokenString := strings.TrimPrefix(auth, "Bearer ")
	request.Header.Add("Authorization", "Bearer "+tokenString)
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "500", Message: http.StatusText(500)})
		log.Println(fmt.Sprintf("ERROR %v", err))
		return
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(response.StatusCode)
		stringStatusCode := strconv.Itoa(response.StatusCode)
		json.NewEncoder(w).Encode(ErrorMessage{Code: stringStatusCode, Message: http.StatusText(response.StatusCode)})
		log.Println("ERROR FROM JWT SERVER: " + stringStatusCode + " (" + urlPath + ") ")
		return
	}

	log.Println("JWT Server Response Status:", response.StatusCode, http.StatusText(response.StatusCode))

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename='%s'", filepath.Base(filePath)))
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	http.ServeFile(w, r, filePath)
}
