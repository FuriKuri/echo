package main

import (
	"fmt"
	"io/ioutil"
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

func respEnv(w http.ResponseWriter, r *http.Request) {
	envs := ""
	for _, pair := range os.Environ() {
		envs = envs + pair + "\n"
	}
	fmt.Fprintf(w, "List of all environment variables:\n%s", envs)
}

func main() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Print("Start server!")
	http.HandleFunc("/hostname", respHostname)
	http.HandleFunc("/env", respEnv)
	http.HandleFunc("/", fallback)
	http.ListenAndServe(":8080", nil)
}
