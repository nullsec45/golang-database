package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_mysql?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	//defer db.Close()
	return db
}
