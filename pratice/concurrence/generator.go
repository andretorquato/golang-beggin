package main

import "time"

func main() {
	channel := write("Hello")
	for i := 0; i < 8; i++ {
		println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- text
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
