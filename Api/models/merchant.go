package models


type Merchant struct{
	Id string `json:"id"`
	Mcc string `json:"mcc"`
	AvgAmount float64 `json:"avg_amount"`
}