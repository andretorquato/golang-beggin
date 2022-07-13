package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func now(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05 PM")
	fmt.Fprintf(w, "<h2>The time is %s </h2>", s)
}

func main() {
	http.HandleFunc("/", now)
	log.Println("executing server...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
