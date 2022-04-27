package main

import "fmt"

func routine(channel chan int) {
	channel <- 1000
	channel <- 2000
	channel <- 3000
	fmt.Println("executed")
	channel <- 4000
}

func main() {
	channel := make(chan int, 2)

	go routine(channel)

	fmt.Println(<-channel)
}
