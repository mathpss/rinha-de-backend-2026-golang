package services

import (
	"slices"
	"api/models"
	"api/utils"
)

func Normalization(payload models.Payload) [14]float64 {
	normalization := utils.GetNormalization()
	listMccRisk := utils.GetMccRisk()
	var terminalIsOnline float64 = 0
	var cardPresent float64 = 0
	var mccRisk float64 = 0.5
	var knownMerchant float64 = 1
	var minuteValue float64 = -1
	var kmValue float64 = -1
	var result [14]float64

	if payload.Terminal.IsOnline{
		terminalIsOnline = 1
	}
	if payload.Terminal.CardPresent{
		cardPresent = 1
	}
	if value, ok:= listMccRisk.MccRisk[payload.Merchant.Mcc]; ok{
		mccRisk = value
	}
	if slices.Contains(payload.Customer.KnownMerchants, payload.Merchant.Id) {
			knownMerchant = 0
		}
 
	if payload.LastTransaction != nil{
		
		minuteValue = limitar((payload.LastTransaction.Timestamp.Hour()*60 +
		 payload.LastTransaction.Timestamp.Minute()),
		normalization.MaxMinutes)

		kmValue = limitar(payload.LastTransaction.KmFromCurrent, normalization.MaxKm)
	
	}

	result[0] = limitar(payload.Transaction.Amount, normalization.Amount)
	result[1] = limitar(payload.Transaction.Installments, normalization.MaxInstallments)
	result[2] = (limitar(payload.Transaction.Amount, payload.Customer.AvgAmount) / normalization.AmountVsAvgRatio)
	result[3] = float64(payload.Transaction.RequestedAt.Hour()) /23
	result[4] = float64(payload.Transaction.RequestedAt.Weekday()) / 6
	result[5] = minuteValue
	result[6] = kmValue
	result[7] = limitar(payload.Terminal.KmFromHome, normalization.MaxKm)
	result[8] = limitar(payload.Customer.TxCount24h, normalization.MaxTxCount24h)
	result[9] = terminalIsOnline
	result[10] = cardPresent
	result[11] = knownMerchant
	result[12] = mccRisk
	result[13] = limitar(payload.Merchant.AvgAmount, normalization.MaxMerchantAvgAmount)

	return result
}

func limitar[T Numeric](a T, b T) float64{
	if a >= b{
		return 1
	}

	return float64(a)/float64(b)
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}