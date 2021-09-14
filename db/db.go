package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetCon() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_test")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
