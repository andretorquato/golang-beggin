package main

import "fmt"

type car struct {
	name  string
	color string
	price float64
}

type tesla struct {
	car
	model string
}

func main() {
	t := tesla{}
	t.name = "Tesla"
	t.color = "Red"
	t.price = 12000.00
	t.model = "Model SX"
	fmt.Println(t)
}
