package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
