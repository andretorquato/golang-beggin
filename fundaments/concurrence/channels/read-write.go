package main

import "fmt"

func main() {
	cha := make(chan string, 1)

	cha <- "first line" // write
	<-cha               // read

	cha <- "using channel"
	fmt.Println(<-cha)
}
