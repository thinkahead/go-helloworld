package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Pi calculator received a request from ", r.RemoteAddr, " for ", r.URL)
	h := w.Header()
	h.Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello world!\n\n")
	fmt.Fprintf(w, "Go version: %s\n", runtime.Version())
}

func main() {
	log.Print("Hello World started.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
