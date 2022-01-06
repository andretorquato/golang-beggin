package main

func swap(p1, p2 int) (first, second int) {
	second = p2
	first = p1
	return
}

func main() {
	x, y := swap(1, 2)
	println(y, x)
}
