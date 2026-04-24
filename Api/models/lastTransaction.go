package models

import "time"

type LastTransaction struct {
	Timestamp time.Time `json:"timestamp"`
	KmFromCurrent float64 `json:"km_from_current"`
}