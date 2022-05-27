package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
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

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User

	db.QueryRow("SELECT id, name from users where id = ?", id).Scan(&u.ID, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func getUsers(w http.ResponseWriter, r *http.Request) {}

func main() {}
