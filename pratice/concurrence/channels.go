package main

import (
	"fmt"
	"time"
)

func main() {
	newChannel := make(chan string)
	go writeN(newChannel, "Hello")

	for {
		message, opened := <-newChannel
		if !opened {
			fmt.Println("--> Channel closed <--")
			break
		}
		fmt.Println(message)
	}

	fmt.Println("Done")
}

func writeN(ch chan string, text string) {
	for i := 1; i < 5; i++ {
		ch <- text
		time.Sleep(time.Second)
	}
	close(ch)
}
