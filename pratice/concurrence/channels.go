package main

import (
	"fmt"
	"time"
)

func main() {
	chanOne := make(chan string)
	chanTwo := make(chan string)
	chanThree := make(chan string)

	go firstChannel(chanOne)
	go secondChannel(chanTwo)
	go thirdChannel(chanThree)

	for message := range chanOne {
		fmt.Println(message)
	}

	for messageTwo := range chanTwo {
		fmt.Println(messageTwo)
	}

	for messageThree := range chanThree {
		fmt.Println(messageThree)
	}

	fmt.Println("Done")
}

func firstChannel(channel chan string) {
	for i := 0; i < 20; i++ {
		channel <- "<-.-> First"
		time.Sleep(time.Second)
	}
	close(channel)
}

func secondChannel(channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- "--- SECOND CHANNEL ---"
		time.Sleep(time.Second * 2)
	}
	close(channel)
}

func thirdChannel(channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- "+++ THIRD CHANNEL +++"
		time.Sleep(time.Second * 3)
	}
	close(channel)
}
