package main

func main() {
	i := 1

	var p *int = &i // &i is the address of i
	println(*p)     // *p is the value of i
	i++
	println(*p)
	*p++
	println(*p)
	println(i)
	println(p) // p is the address of i

}
