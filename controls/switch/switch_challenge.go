package main

import "fmt"

func getNoteForConcept(note float64) string {
	switch {
	case note >= 9:
		return "A"
	case note < 9 && note >= 7:
		return "B"
	case note < 7 && note >= 5:
		return "C"
	case note < 5 && note >= 3:
		return "D"
	case note < 3:
		return "F"
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(getNoteForConcept(10))
}
