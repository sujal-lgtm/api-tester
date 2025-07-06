package main

import (
	"log"
	"net/http"
	"time"
	"testapi/internal/handler"
)

var WriteTimeout = 1 * time.Second

func main() {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:         ":6970",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	router.HandleFunc("/test", handler.HandleTest)

	log.Println("🚀 Server starting on port :6970")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
