package main

import (
	"ProjectMIRRO/backend/gateway/handlers"
	"ProjectMIRRO/backend/gateway/models/users"
	"ProjectMIRRO/backend/gateway/sessions"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

//main is the main entry point for the server
func main() {

	sessionID := os.Getenv("SESSIONKEY")
	redisAddr := os.Getenv("REDISADDR")
	TLSKEY := os.Getenv("TLSKEY")
	TLSCERT := os.Getenv("TLSCERT")
	addr := ":443"

	DSN := os.Getenv("DSN")
	if len(redisAddr) == 0 {
		redisAddr = "127.0.0.1:6379"
	}
	redisDB := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	sessionStore := sessions.NewRedisStore(redisDB, time.Hour)

	userStore, err := users.NewMySQLStore(DSN)
	if err != nil {
		log.Printf("Unable to open database mysql %v", err)
	}

	contextHandler := &handlers.ContextHandler{
		SessionID:    sessionID,
		SessionStore: sessionStore,
		UserStore:    userStore,
	}

	mux := http.NewServeMux()
	log.Printf("server is listening at %s...", addr)
	mux.HandleFunc("/users", contextHandler.UsersHandler)
	mux.HandleFunc("/users/", contextHandler.SpecificUserHandler)
	mux.HandleFunc("/sessions", contextHandler.SessionsHandler)
	mux.HandleFunc("/sessions/", contextHandler.SpecificSessionHandler)
	wrappedMux := handlers.NewHeaderHandler(mux)
	log.Fatal(http.ListenAndServeTLS(addr, TLSCERT, TLSKEY, wrappedMux))
}
