package math

import (
	"fmt"
	"strconv"
)

func Average(values ...float64) float64 {
	var total float64
	for _, value := range values {
		total += value
	}
	average := total / float64(len(values))
	roundedAverage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	return roundedAverage
}
