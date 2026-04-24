package models

import "time"

type Transaction struct {
	Amount       float64 `json:"amount"`
	Installments int `json:"installments"`
	RequestedAt  time.Time `json:"requested_at"`
}