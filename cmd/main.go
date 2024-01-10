package main

import (
	"log"
	"net/http"
	"time"

	router "github.com/faustocalvinio/go_todo_api_mongodb/routes/routes.go"
)

func main() {
    srv := &http.Server{
		Addr:         ":9000",
		Handler:      router.Router(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
    log.Fatal(srv.ListenAndServe())
}