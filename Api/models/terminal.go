package models


type Terminal struct{
	IsOnline bool `json:"is_online"`
	CardPresent bool `json:"card_present"`
	KmFromHome float64 `json:"km_from_home"`
}