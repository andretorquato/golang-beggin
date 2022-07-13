package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")

	stmt.Exec(1, "John")
	stmt.Exec(2, "Jane")
	_, err = stmt.Exec(1, "Joãoe")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
