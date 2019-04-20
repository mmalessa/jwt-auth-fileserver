package main

import (
	"net/http"
	"html"
	"io/ioutil"
	"time"
	"bytes"
	"log"
	"strings"
	"encoding/json"
)

type ErrorMessage struct{
	Code string
	Message string
}

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
		log.Println("ERROR 401: /")
		return

		// fmt.Fprintf(w, "Welcome message")
		// fmt.Printf("%s %s %s \n", r.Method, r.URL, r.Proto)
		//Iterate over all header fields
		// for k, v := range r.Header {
		// 	fmt.Printf("Header field %q, Value %q\n", k, v)
		// }
		// fmt.Printf("Host = %q\n", r.Host)
		// fmt.Printf("RemoteAddr= %q\n", r.RemoteAddr)
		// fmt.Printf("\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
	})

	http.HandleFunc("/files/", func (w http.ResponseWriter, r *http.Request) {

		urlpath := html.EscapeString(r.URL.Path)
		filename := strings.TrimPrefix(urlpath, "/files/")
		path := string("files/" + filename)

		auth := r.Header.Get("Authorization")
		if (auth == "") {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
			log.Println("NO TOKEN (" + urlpath + ")")
			return
		}

		// FIXME - here test JWT
		// token = strings.TrimPrefix(auth, "Bearer ")
		// bearer := r.Header.Bearer (??)
		
		data, err := ioutil.ReadFile(path)
        if err != nil { 
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(ErrorMessage{Code: "404", Message: http.StatusText(404)})
			log.Println("NOT FOUND (" + urlpath + ")")
			return
		}

		log.Println("DOWNLOAD: " + urlpath)

		http.ServeContent(w, r, path, time.Now(), bytes.NewReader(data))
	})

	http.ListenAndServe(":8080", nil)
}