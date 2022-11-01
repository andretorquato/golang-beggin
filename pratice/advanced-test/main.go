package main

import (
	"fmt"
)

func main() {
	sumPoints := calculate.Point{X: 10, Y: 20}.Sum()
	fmt.Println(sumPoints)
}
