package main

import "fmt"

func main() {

	fmt.Print("Sem ")
	fmt.Print("quebra linha \n")

	fmt.Println("Quebra linha")
	fmt.Println("automaticamente")

	x := 3.12345

	xs := fmt.Sprint(x)
	fmt.Println("Concatened value " + xs)
	fmt.Printf("The value of x is %f\n", x)
	fmt.Printf("The value of x is %.2f\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("%d %f %.1f %t %s\n", a, b, b, c, d)
}
