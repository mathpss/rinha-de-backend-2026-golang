package models


type Payload struct{
	Id string `json:"id"`
	Customer Customer `json:"customer"`
	Transaction Transaction `json:"transaction"`
	Merchant Merchant `json:"merchant"`
	Terminal Terminal `json:"terminal"`
	LastTransaction  *LastTransaction `json:"last_transaction"`
}