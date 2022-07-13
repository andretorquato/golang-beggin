package main

import (
	"fmt"
	"time"
)

func simplePassDataToChannel(channel chan int) {
	time.Sleep(time.Second)
	channel <- 1000
	fmt.Println("Sended before read")
}

func main() {
	channel := make(chan int)

	go simplePassDataToChannel(channel)

	fmt.Println(<-channel)
	fmt.Println("Read")

	fmt.Println(<-channel) // deadlock
}
