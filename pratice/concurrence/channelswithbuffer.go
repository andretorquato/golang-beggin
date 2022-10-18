package main

func main() {
	channel := make(chan string, 2)

	channel <- "First"
	channel <- "Second"

	responseOne := <-channel
	responseTwo := <-channel

	println(responseOne)
	println(responseTwo)
}
