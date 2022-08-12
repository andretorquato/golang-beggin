package main

import "time"

func main() {
	go write("Hello -- ")
	write("-- world")
}

func write(text string) {
	for {
		println(text)
		time.Sleep(time.Second)
	}
}
