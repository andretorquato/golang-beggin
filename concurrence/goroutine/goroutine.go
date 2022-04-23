package main

import (
	"fmt"
	"time"
)

func execute(cmd string, quantity int) {
	for t := 0; t < quantity; t++ {
		time.Sleep(time.Second)
		fmt.Println(cmd, ":", t)
	}
}

func main() {
	const DEFAULT_QUANTITY = 3
	// execute("Hello world no thread 2", 3)
	// execute("Hello world no thread 1", 3)

	// with goroutines
	go execute("Hello world no thread 1", DEFAULT_QUANTITY)
	go execute("Hello world no thread 2", DEFAULT_QUANTITY)

	time.Sleep(time.Second*DEFAULT_QUANTITY + 4)
}
