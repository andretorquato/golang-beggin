package main

import "fmt"

func factorial(num int) (int, error) {
	switch {
	case num < 0:
		return -1, fmt.Errorf("Invalid number: %d", num)
	case num == 0:
		return 1, nil
	default:
		prodNum, _ := factorial(num - 1)
		return num * prodNum, nil
	}
}

func simpleFactorial(num int) int {
	if num <= 0 {
		return 1
	}
	return num * simpleFactorial(num-1)
}

func main() {
	result, _ := factorial(12)
	fmt.Println(result)
	_, error := factorial(-2)

	if error != nil {
		fmt.Println(error)
	}

	result2 := simpleFactorial(12)
	fmt.Println(result2)
}
