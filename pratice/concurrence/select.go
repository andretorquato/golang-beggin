package main

import "time"

func main() {
	channelOne, channelTwo := make(chan string), make(chan string)

	go func() {
		for {
			channelOne <- "First"
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			channelTwo <- "Second"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case messageOne := <-channelOne:
			println(messageOne)
		case messageTwo := <-channelTwo:
			println(messageTwo)
		}
	}
}
