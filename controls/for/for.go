package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	for {
		fmt.Println("Hello")
		time.Sleep(time.Second)
	}
}
