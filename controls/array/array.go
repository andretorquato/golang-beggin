package main

import "fmt"

func main() {
	var notes [3]float64

	notes[0], notes[1], notes[2] = 4.5, 9.7, 10.0

	total := 0.0

	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	average := total / float64(len(notes))
	fmt.Printf("Average: %.2f\n", average)
}
