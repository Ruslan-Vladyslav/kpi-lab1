package main

import (
	"kpi-lab1/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/time", server.TimeHandler)

	log.Println("Server is running on http://localhost:8795")
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		log.Fatal(err)
	}
}
