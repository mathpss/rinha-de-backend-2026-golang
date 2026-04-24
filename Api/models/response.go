package models

type Response struct{
	Approved bool `json:"approved"`
	FraudScore float64 `json:"fraud_score"`
}