package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, smtError := db.Prepare("insert into users(name) values(?)")

	if smtError != nil {
		panic(err)
	}

	stmt.Exec("Jane")
	stmt.Exec("Joe")

	res, _ := stmt.Exec("John")

	id, _ := res.LastInsertId()
	fmt.Println(id)

}
