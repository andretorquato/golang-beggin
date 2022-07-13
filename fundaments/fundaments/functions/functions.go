package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func show(a int) {
	fmt.Println(a)
}

func main() {
	result := sum(1, 5)
	show(result)
}
