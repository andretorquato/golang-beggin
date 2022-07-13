package main

import "fmt"

type printable interface {
	toString() string
}

type people struct {
	name     string
	lastname string
}

type student struct {
	people
	registration int
}

func (p people) toString() string {
	return p.name + " " + p.lastname
}

func (s student) toString() string {
	return fmt.Sprintf("name: %s, lastname: %s, registration: %d", s.name, s.lastname, s.registration)
}

func print(p printable) {
	fmt.Println(p.toString())
}

func main() {
	var any printable = student{people{"John", "Doe"}, 12345}

	print(any)

	person := people{"John", "Doe"}
	print(person)
}
