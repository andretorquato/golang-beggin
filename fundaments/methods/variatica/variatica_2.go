package main

import "fmt"

// with slice
func printApproved(names ...string) {
	for i, name := range names {
		fmt.Printf("%d) %s\n", i+1, name)
	}
}

func main() {
	approved := []string{"John", "Paul", "George", "Ringo", "Garcia"}
	printApproved(approved...)
}
