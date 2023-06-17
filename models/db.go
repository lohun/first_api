package models

import (
	"database/sql"
	"log"
	"time"
	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB  = connection()

func connection() *sql.DB {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "pm",
		AllowNativePasswords: true,
	}
	Db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	return Db
}