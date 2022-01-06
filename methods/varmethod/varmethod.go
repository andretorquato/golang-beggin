package main

var sum = func(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	println(sum(1, 2))

	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}
	println(sub(1, 2))
}
