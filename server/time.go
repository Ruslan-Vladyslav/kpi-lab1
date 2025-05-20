package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	// Use RFC3339 format
	currentTime := time.Now().Format(time.RFC3339)

	// Log req
	log.Printf("GET /time -> %s\n", currentTime)
	log.Printf("Request from %s to %s", r.RemoteAddr, r.URL.Path)

	// Get current time
	response := TimeResponse{
		Time: currentTime,
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)

	// If error -- log
	if err != nil {
		log.Fatal(err)
	}
}
