package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root@tcp(localhost)/go_graphql_test")
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
