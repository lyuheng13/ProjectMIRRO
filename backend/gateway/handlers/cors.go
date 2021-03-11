package handlers

import (
	"log"
	"net/http"
	"time"
)

type HeaderHandler struct {
	Handler http.Handler
}

func (hh *HeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Max-Age", "600")
	start := time.Now()
	hh.Handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))

}

// NewHeaderHandler wraps the http.Handler into middleware structure
func NewHeaderHandler(handlerToWrap http.Handler) *HeaderHandler {
	return &HeaderHandler{handlerToWrap}
}
