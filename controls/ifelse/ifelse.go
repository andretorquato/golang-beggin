package main

func showResult(note float64) {
	if note >= 7 {
		println("You passed!")
	} else {
		println("You failed!")
	}
}

func main() {
	showResult(7.0)
	showResult(5.0)
}
