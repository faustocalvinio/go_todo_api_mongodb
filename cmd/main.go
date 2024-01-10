package main

import (
	"log"
	"net/http"
	"time"
	"todo_api_mongodb/routes"
)

func main() {
    srv := &http.Server{
		Addr:         ":9000",
		Handler:      routes.Router(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
    log.Println("Server is running on http://localhost:9000")
	log.Fatal(srv.ListenAndServe())
}