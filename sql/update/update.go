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

	stmt, _ := db.Prepare("UPDATE users set name = ? where id = ?")

	stmt.Exec("Clark Quiston", 1)
	stmt.Exec("Jane Maclaren", 2)
	stmt.Exec("Ronaldo Montblanc", 3)

	stmt2, _ := db.Prepare("DELETE FROM users where id = ?")

	stmt2.Exec(1)
}
