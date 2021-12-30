package main

import "fmt"

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float32, float64:
		return "float"
	case func():
		return "func"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(getType(10))
	fmt.Println(getType("Hello"))
	fmt.Println(getType(10.5))
	fmt.Println(getType(func() {}))
}
