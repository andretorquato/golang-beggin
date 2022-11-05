package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

func main() {
	p := Person{
		Name:   "John",
		Age:    30,
		Salary: 5000.0,
	}

	// Marshal
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
	fmt.Println(bytes.NewBuffer(b))

	p2 := map[string]string{
		"name":   "John",
		"age":    "30",
		"salary": "5000.0",
	}

	// Marshal
	b2, err := json.Marshal(p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b2)
	fmt.Println(bytes.NewBuffer(b2))

}
