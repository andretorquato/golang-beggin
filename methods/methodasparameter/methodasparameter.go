package main

func productOfNumber(a, b int) int {
	return a * b
}

func exec(method func(int, int) int, a int, b int) int {
	return method(a, b)
}

func main() {
	println(exec(productOfNumber, 2, 3))
}
