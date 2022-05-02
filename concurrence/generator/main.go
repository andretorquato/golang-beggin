package main

import (
	"fmt"

	"github.com/andretorquato/golang-beggin/reusable/html"
)

func main() {
	t1 := html.Title("https://www.baidu.com", "https://www.google.com")
	t2 := html.Title("https://www.gmail.com", "https://www.amazon.com")

	fmt.Printf("%s | %s\n", <-t1, <-t2)
	fmt.Printf("%s | %s\n", <-t1, <-t2)
}
