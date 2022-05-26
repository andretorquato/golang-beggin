package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/user/")
	id, _ := strconv.Atoi(sid)
	switch {
	case r.Method == "GET" && id > 0:
		getUser(w, r, id)
	case r.Method == "GET":
		getUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("not found")
	}
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {}

func getUsers(w http.ResponseWriter, r *http.Request) {}

func main() {}
