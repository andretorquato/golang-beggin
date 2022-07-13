package main

import "fmt"

func speak(user string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			channel <- user + " says " + string(i)
		}
	}()
	return channel
}

func multiplex(enter1, enter2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case t1 := <-enter1:
				channel <- t1
			case t2 := <-enter2:
				channel <- t2
			}
		}
	}()
	return channel
}
func main() {
	cc := multiplex(speak("John"), speak("Mary"))
	fmt.Println(<-cc, <-cc)
	fmt.Println(<-cc, <-cc)
	fmt.Println(<-cc, <-cc)
}
