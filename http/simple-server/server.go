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
	fmt.Println(id)
}

func main() {}
