package models

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB  = connection()

func connection() *sql.DB {
	Db, err := sql.Open("mysql", "root@locahost")
	if err != nil {
		log.Fatalln(err)
	}
	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	return Db
}