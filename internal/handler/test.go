package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"testapi/internal/models"
	"testapi/internal/tester"
)

func HandleTest(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Received /test request")

	var req models.TestRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println("✅ Parsed request:", req)

	var result interface{}
	var err error

	switch req.Type {
	case "performance":
		log.Println("🚀 Running performance test...")
		result, err = tester.RunPerformanceTest(req)

	case "load":
		log.Println("🔥 Running load test...")
		result, err = tester.RunLoadTest(req)

	default:
		log.Println("❌ Invalid test type:", req.Type)
		http.Error(w, "Invalid test type", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Println("❌ Test failed:", err)
		http.Error(w, "Test Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Marshal the struct result to JSON
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode result", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Sending result")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
