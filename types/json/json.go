package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	newProduct := product{"Mouse", 10.0, []string{"computer", "mouse"}}
	fmt.Println(newProduct)
	productJson, _ := json.Marshal(newProduct)
	fmt.Println(string(productJson))

	var product2 product
	json.Unmarshal(productJson, &product2)
	fmt.Println(product2.Tags[1])
}
