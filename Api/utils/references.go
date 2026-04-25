package utils

import (
	_ "embed"
	"encoding/json"
)

//go:embed references.json
var fileData []byte

type References struct{
	Vector [14]float64 `json:"vector"`
	Label string `json:"label"`
}

var cached []References

func init(){
	err := json.Unmarshal(fileData, &cached)
	if err !=nil {
		panic(err)
	}
}

func GetReferences() []References{
	return cached
}