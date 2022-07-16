package main

import "fmt"

func main() {
	// explicita
	var var1 string = "Variable 1"
	fmt.Println(var1)

	// inferência de tipo
	var2 := "Variable 2"
	fmt.Println(var2)

	// multiple declaration
	var (
		var3 string = "Variable 3"
		var4 string = "Variable 4"
	)
	fmt.Println(var3, var4)
	const constant1 = "Constant 1"
	fmt.Println(constant1)

	var5, var6 := "Variable 5", "Variable 6"
	fmt.Println(var5, var6)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}
