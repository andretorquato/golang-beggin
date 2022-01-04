package main

import "runtime/debug"

func method3() {
	debug.PrintStack()
}
func method2() {
	method3()
}
func method1() {
	method2()
}

func main() {
	method1()
}
