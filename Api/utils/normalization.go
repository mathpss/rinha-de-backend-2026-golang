package utils

type Normalization struct{
	Amount float64 `json:"max_amount"` 
	MaxInstallments int `json:"max_installments"`
	AmountVsAvgRatio float64 `json:"amount_vs_avg_ratio"`
	MaxMinutes int `json:"max_minutes"` // caso de algum erro de calculo olhar aqui
	MaxKm float64 `json:"max_km"`
	MaxTxCount24h int `json:"max_tx_count_24h"`
	MaxMerchantAvgAmount float64 `json:"max_merchant_avg_amount"`
}

func GetNormalization() *Normalization{
	return &Normalization{
		Amount: 10000,
		MaxInstallments: 12,
		AmountVsAvgRatio: 10,
		MaxMinutes: 1440,
		MaxKm: 1000,
		MaxTxCount24h: 20,
		MaxMerchantAvgAmount: 10000,
	}
}