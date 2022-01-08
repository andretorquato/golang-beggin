package main

import "fmt"

func getApprovedValue(value int) int {
	defer fmt.Printf("end\n")
	if value > 100 {
		fmt.Println("value is greater than 100")
	}
	return value
}
func main() {
	fmt.Println(getApprovedValue(1000))
}
