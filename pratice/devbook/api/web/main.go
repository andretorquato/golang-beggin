package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func main() {
	r := router.Generate()
	utils.LoadTemplate()

	fmt.Println("Running web application on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
