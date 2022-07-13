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

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
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

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(u.Name)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	u.ID = int(lastId)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name from users")

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
