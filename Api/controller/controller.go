package controller

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
)

func ReadyHandle(w http.ResponseWriter, r *http.Request) {

}

func FraudScoreHandle(w http.ResponseWriter, r *http.Request){
	var payload models.Payload
	err:= json.NewDecoder(r.Body).Decode(&payload)	
	if err != nil{
		http.Error(w, "Erro ao ler o json "+ err.Error(), http.StatusBadRequest)
		return
	}

	var normalized [14]float64 = services.Normalization(payload)

	result := services.EuclidianTop5(normalized)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
