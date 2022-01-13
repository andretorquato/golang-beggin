package main

import "fmt"

type note float64

func (n note) entered(init, end float64) bool {
	return float64(n) >= init && float64(n) <= end
}

func noteForConcept(n note) string {
	switch {
	case n.entered(9.0, 10):
		return "A"
	case n.entered(7.0, 8):
		return "B"
	case n.entered(5.0, 6):
		return "C"
	case n.entered(3.0, 4):
		return "D"
	default:
		return "F"
	}
}

func main() {
	fmt.Println(noteForConcept(1.0))
}
