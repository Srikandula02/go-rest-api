package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect() {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "oct@0210",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "employee",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("DB Connected")
}

func GetDB() *sql.DB {
	return DB
}
