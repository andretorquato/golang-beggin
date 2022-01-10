package main

import (
	"fmt"
	"strings"
)

type people struct {
	name     string
	lastname string
}

func (p people) getCompletedName() string {
	return p.name + " " + p.lastname
}

func (p *people) setCompletedName(name string) {
	words := strings.Split(name, " ")
	p.name = words[0]
	p.lastname = words[1]
}

func main() {
	people := people{"Torquato", "André"}
	fmt.Println(people.getCompletedName())

	people.setCompletedName("André Torquato")
	fmt.Println(people.getCompletedName())
}
