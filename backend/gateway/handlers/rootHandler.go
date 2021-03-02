package handlers

import (
	"fmt"
	"net/http"
)

const headerCORS = "Access-Control-Allow-Origin"
const corsAnyOrigin = "*"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	w.Header().Add(headerCORS, corsAnyOrigin)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}
