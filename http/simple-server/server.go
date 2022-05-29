package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", UserHandler)
	http.HandleFunc("/users/add", AddUserHandler)
	log.Println("Running on port 30000")
	http.ListenAndServe(":3000", nil)
}
