package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {
	fmt.Println("Running web application")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":8080", r))
}
