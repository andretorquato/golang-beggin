package main

import (
	"fmt"
	"time"
)

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primes(total int, channel chan int) {
	initial := 2
	for i := 0; i < total; i++ {
		for prime := initial; ; prime++ {
			if isPrime(prime) {
				channel <- prime
				initial = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(channel)
}

func main() {
	channel := make(chan int, 15)

	go primes(cap(channel), channel)

	for prime := range channel {
		fmt.Printf("%d ", prime)
	}
}
