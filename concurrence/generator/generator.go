package main

import "fmt"

func title(urls ...string) {
	for _, url := range urls {
		fmt.Println(url)
	}

}

func main() {
	title("Google", "Baidu", "Runoob")
}
