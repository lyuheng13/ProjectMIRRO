package main

import (
	"log"
	"net/http"
	"os"
)

// This is the main function for the gateway/firewall
func main() {
	addr := os.Getenv("ADDR")
	TLSKEY := os.Getenv("TLSKEY")
	TLSCERT := os.Getenv("TLSCERT")
	if len(addr) == 0 {
		addr = ":443"
	}

	// Route requests to different handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	log.Printf("now server is listening")
	log.Fatal(http.ListenAndServeTLS(addr, TLSKEY, TLSCERT, mux))
}
