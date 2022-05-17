package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Begin()
}
