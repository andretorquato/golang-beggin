package main

func noteForConcept(note float64) string {
	if note >= 8 {
		return "A"
	} else if note < 8 && note >= 7 {
		return "B"
	} else if note < 7 && note >= 6 {
		return "C"
	} else if note < 7 && note >= 4 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	println(noteForConcept(7.0))
	println(noteForConcept(5.0))
}
