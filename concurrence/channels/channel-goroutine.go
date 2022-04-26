package main

import (
	"fmt"
	"time"
)

func multipleOfNumber(value chan int, base int) {
	time.Sleep(time.Second)
	value <- 3 * base

	time.Sleep(time.Second)
	value <- 4 * base

	time.Sleep(time.Second)
	value <- 5 * base
}

func main() {
	value := make(chan int)

	go multipleOfNumber(value, 2)
	fmt.Println("First execution")
	value1, value2 := <-value, <-value
	fmt.Println("Second execution")

	fmt.Println(value1, value2)
	// time.Sleep(time.Second)

	fmt.Println(<-value)
}
