package controller

import (
	"api/models"
	"encoding/json"
	"net/http"
)

func ReadyHandle(w http.ResponseWriter, r *http.Request) {

}

func FraudScoreHandle(w http.ResponseWriter, r *http.Request){
	var payload models.Payload
	var teste models.Response = models.Response{Approved: true, FraudScore: 1.0}
	err:= json.NewDecoder(r.Body).Decode(&payload)

	if err != nil{
		http.Error(w, "Erro ao ler o json "+ err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teste)
}
