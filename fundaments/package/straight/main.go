package main

import "fmt"

func main() {
	p1 := Point{1, 0}
	p2 := Point{0, 1}

	fmt.Println(p1, p2)

	fmt.Println(cateto(p1, p2))
	fmt.Println(Distance(p1, p2))

}
