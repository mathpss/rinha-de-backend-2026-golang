package main

import (
	"api/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /ready", controller.ReadyHandle)

	http.HandleFunc("POST /fraud-score", controller.FraudScoreHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
