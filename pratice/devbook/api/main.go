package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Print("Listening..")
	r := router.Start()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
