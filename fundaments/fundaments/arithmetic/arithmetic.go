package main

import "fmt"

func main() {
	a := 4
	b := 2

	fmt.Println("Sum:", a+b)
	fmt.Println("Difference:", a-b)
	fmt.Println("Product:", a*b)
	fmt.Println("Quotient:", a/b)
	fmt.Println("Remainder:", a%b)

	// bitwise operators
	fmt.Println("Bitwise AND:", a&b) // 11 & 10 = 10
	fmt.Println("Bitwise OR:", a|b)  // 11 | 10 = 11
	fmt.Println("Bitwise XOR:", a^b) // 11 ^ 10 = 01

}
