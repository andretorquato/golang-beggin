package main

import "fmt"

func getNoteForConcept(n float64) string {
	var note = n
	switch note {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "F"
	default:
		return "Nota inválida"
	}

}

func main() {
	fmt.Print(getNoteForConcept(10))
}
