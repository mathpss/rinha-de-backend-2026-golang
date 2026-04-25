package services

import (
	"api/models"
	"api/utils"
	"sort"
)

type checkedVector struct {
	vector [14]float64
	label  string
	distace float64
}

func EuclidianTop5(normalized [14]float64) models.Response {
	var result []checkedVector
	references := utils.GetReferences()
	
	for _, value := range references {
		distance := euclidianCalc(normalized, value.Vector)

		result = append(result, checkedVector{
			vector: value.Vector,
			label: value.Label,
			distace: distance,
		})
	}

	sort.Slice(result, func(i,j int) bool{
		return result[i].distace < result[j].distace
	})

	if len(result) > 5 {
		return scoreCalculate(result[:5])
	}
	return scoreCalculate(result)
}

func euclidianCalc(v1, v2 [14]float64) float64 {
	var sum float64
	for i := range 14 {
		diff := v1[i] - v2[i]
		sum += diff * diff
	}
	return sum
}

func scoreCalculate(checker []checkedVector) models.Response{
	var counterFraud float64

	for _, value := range checker{
		if value.label == "fraud" {
			counterFraud ++
		}
	}

	return models.Response{
		Approved: (counterFraud / 5) < 0.6 ,
		FraudScore: counterFraud / 5,
	}

}