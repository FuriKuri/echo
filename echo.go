package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func fallback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func respHostname(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "My hostname is %s!", hostname)
}

func main() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Print("Start server!")
	http.HandleFunc("/hostname", respHostname)
	http.HandleFunc("/", fallback)
	http.ListenAndServe(":8080", nil)
}
