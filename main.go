package main

import (
	"log"
	"net/http"
	"kpi-lab1/server"
)

func main() {
	http.HandleFunc("/time", server.TimeHandler)

	log.Println("Сервер запущен на http://localhost:8795")
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		log.Fatal(err)
	}
}
