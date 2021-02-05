package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const headerCORS = "Access-Control-Allow-Origin"
const corsAnyOrigin = "*"

func rootHandler(w http.ResponseWriter, read *http.Request) {
	name := read.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	w.Header().Add(headerCORS, corsAnyOrigin)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func main() {
	addr := os.Getenv("ADDR")

	if len(addr) == 0 {
		addr = ":80"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	log.Printf("now server is listening")
	log.Fatal(http.ListenAndServe(addr, mux))
}
