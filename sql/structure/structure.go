package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS golangcourse")
	exec(db, "USE golangcourse")
	exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `
		CREATE TABLE users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

}
