package main

import (
	"fmt"
	"time"
)

func main() {

	// 'André' == 'André' = true
	// 1 > 2 = false
	// 1 < 2 = true
	// 1 >= 2 = false
	// 1 <= 2 = true
	// 1 != 2 = true
	// 1 == 2 = false

	date1 := time.Unix(0, 0)
	date2 := time.Unix(0, 0)

	fmt.Println(date1 == date2)
	fmt.Println(date1.Equal(date2))

	type Person struct {
		name string
	}

	p1 := Person{name: "André"}
	p2 := Person{name: "André"}
	fmt.Println(p1 == p2)
}
