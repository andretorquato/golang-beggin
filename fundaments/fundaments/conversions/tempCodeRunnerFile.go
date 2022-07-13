package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Print("Conversão referente ao valor na tabela ASCII ")
	fmt.Print(string(97), "\n")

	// string to int
	n, _ := strconv.Atoi("123")
	fmt.Println(n - 122)

	// boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}

}
