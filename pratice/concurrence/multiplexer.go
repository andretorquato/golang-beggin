package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplexer(write("Hello"), write("World"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channels ...<-chan string) <-chan string {
	merged := make(chan string)
	for _, channel := range channels {
		go func(channel <-chan string) {
			for {
				merged <- <-channel
			}
		}(channel)
	}
	return merged
}

func write(message string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("%s %s", message, time.Now().Format("15:04:05"))
			time.Sleep(time.Second)
		}
	}()

	return channel
}
