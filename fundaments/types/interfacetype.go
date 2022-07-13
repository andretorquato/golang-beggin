package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	thing = course{"Go"}
	fmt.Println(thing)
	thing = 1
	fmt.Println(thing)
	type dynamicType struct {
		a string
		b string
	}
	thing = dynamicType{a: "hello", b: "world"}
	fmt.Println(thing)

	type dynamic interface{}
	var thing2 dynamic = "Hello"
	fmt.Println(thing2)
}
