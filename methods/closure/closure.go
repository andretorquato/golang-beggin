package main

func closure() func() {
	i := 10
	return func() {
		println(i)
	}
}

func main() {
	x := 30
	println(x)
	printX := closure()
	printX()
}
