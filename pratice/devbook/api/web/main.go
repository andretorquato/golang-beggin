package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/src/utils"
)

func main() {
	config.Load()
	cookies.Configure()

	r := router.Generate()
	utils.LoadTemplate()

	fmt.Printf("Running web application on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
