package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Average of [1, 2, 3, 4, 5]: %.2f\n", average(1, 2, 3, 4, 5))
}
