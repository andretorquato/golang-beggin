package main

import (
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// init var here
	if n := randomNumber(); n > 5 {
		println("You won!")
	} else {
		println("You lost!")
	}
}
