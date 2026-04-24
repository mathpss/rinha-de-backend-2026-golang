package models


type Customer struct{
	AvgAmount float64 `json:"avg_amount"`
	TxCount24h int `json:"tx_count_24h"`
	KnownMerchants []string `json:"known_merchants"`
}