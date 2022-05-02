package main

import (
	"fmt"

	"github.com/andretorquato/golang-beggin/reusable/html"
)

func forward(origin <-chan string, destination chan string) <-chan string {

	for {
		destination <- <-origin
	}
}

func multiplex(entries ...<-chan string) <-chan string {
	channel := make(chan string)
	for _, entry := range entries {
		go forward(entry, channel)
	}
	return channel
}

func main() {
	channel := multiplex(
		html.Title("https://www.baidu.com", "https://www.google.com"),
		// html.Title("https://fluenceconsultoria.com/", "https://www.amazon.com"),
	)
	fmt.Printf("%s | %s\n", <-channel, <-channel)
	fmt.Printf("%s | %s\n", <-channel, <-channel)
}
