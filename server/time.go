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
	currentTime := time.Now().Format(time.RFC3339)

	log.Printf("GET /time -> %s\n", currentTime)

	log.Printf("Request from %s to %s", r.RemoteAddr, r.URL.Path)

	response := TimeResponse{
		Time: currentTime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
