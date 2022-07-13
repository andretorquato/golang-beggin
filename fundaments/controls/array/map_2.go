package main

import "fmt"

func main() {
	workers := map[string]float64{
		"Andre": 199.99,
		"João":  199.99,
		"Maria": 199.99,
	}

	fmt.Println(workers)
	workers["Tim"] = 299.99

	for name, salary := range workers {
		fmt.Printf("%s (salary: %.2f)\n", name, salary)
	}

	delete(workers, "Maria")
	delete(workers, "Nonexistent")
}
