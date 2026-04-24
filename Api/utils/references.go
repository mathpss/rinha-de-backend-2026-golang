package utils

import (
	_ "embed"
	"encoding/json"
)

//go:embed references.json
var fileData []byte

type References struct{
	Vector []float64 `json:"vector"`
	Label string `json:"label"`

}

func GetReferences() []References{
	var list []References

	err := json.Unmarshal(fileData, &list)

	if err !=nil {
		panic(err)
	}

	return list
}