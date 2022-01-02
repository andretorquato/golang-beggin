package main

import "fmt"

func main() {
	phoneList := map[string]map[string]string{
		"A": {
			"Andre": "123456789",
			"Arian": "987654321",
		},
		"B": {
			"Bruno":  "3243456789",
			"Baskar": "3213234567",
		},
		"C": {
			"Caio":   "987654321",
			"Carlos": "9876543213",
		},
	}

	for letter := range phoneList {
		for name, phone := range phoneList[letter] {
			fmt.Printf("%s - %s\n", name, phone)
		}
	}

}
