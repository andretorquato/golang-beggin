package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Listening..")
	r := router.Start()

	log.Fatal(http.ListenAndServe(":5000", r))
}
