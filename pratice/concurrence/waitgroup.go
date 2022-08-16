package main

import (
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("First")
		waitGroup.Done()
	}()

	go func() {
		write("Second")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 4; i++ {
		println(text)
		time.Sleep(time.Second)
	}
}
