package main

import "fmt"

type item struct {
	id       int
	quantity int
	price    float64
}

type order struct {
	id    int
	items []item
}

func (o order) getTotalPrice() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {
	order := order{
		id: 1131231,
		items: []item{
			{id: 1, quantity: 2, price: 100},
			{id: 2, quantity: 3, price: 200},
		},
	}

	fmt.Printf("Total is: $%.2f\n", order.getTotalPrice())
}
