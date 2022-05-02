package main

import (
	"fmt"

	"../reusable/html"
)

func main() {
	t1 := html.Ti("https://www.baidu.com", "https://www.google.com")
	t2 := title("https://www.gmail.com", "https://www.amazon.com")

	fmt.Printf("%s | %s\n", <-t1, <-t2)
	fmt.Printf("%s | %s\n", <-t1, <-t2)
}
