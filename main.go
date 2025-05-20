package main

import (
	"kpi-lab1/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/time", server.TimeHandler)

	// Log
	log.Println("Server is running on http://localhost:8795")

	// Start
	err := http.ListenAndServe(":8795", nil)

	// If Error -- log
	if err != nil {
		log.Fatal(err)
	}
}
