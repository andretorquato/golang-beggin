package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golangcourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM users where id > ?", 0)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u users
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			panic(err)
		}
		fmt.Println(u)
	}

}
