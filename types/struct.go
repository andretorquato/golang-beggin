package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) getPriceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	product1 := product{
		name:     "laptop",
		price:    1000,
		discount: 0.09,
	}

	fmt.Println(product1.getPriceWithDiscount())

	product2 := product{"mobile", 2000, 0.05}

	fmt.Println(product2)
}
