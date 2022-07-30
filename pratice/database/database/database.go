package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", "root:password@/devbook?charset=utf8&parseTime=True&loc=Local")

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
