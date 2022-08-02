package main

import (
	"database-pratice/src/server"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
