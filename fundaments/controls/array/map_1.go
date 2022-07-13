package main

import "fmt"

func main() {
	peoples := make(map[int]string)

	peoples[1232123] = "João"
	peoples[543123] = "Fracis"
	peoples[564321] = "Luis"

	fmt.Println(peoples)

	for id, name := range peoples {
		fmt.Printf("%s (id: %d)\n", name, id)

	}

	fmt.Println(peoples[564321])

	delete(peoples, 564321)
	fmt.Println(peoples)
}
