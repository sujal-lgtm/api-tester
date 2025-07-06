package main

import (
	"log"
	"net/http"
	"time"
	"testapi/internal/handler"
	"os"
)

var WriteTimeout = 1 * time.Second

func main() {
	router := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	router.HandleFunc("/test", handler.HandleTest)

	log.Println("🚀 Server starting on port :6969")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
