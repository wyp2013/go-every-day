package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql(conn string, maxOpen, maxIdle int) (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	if maxOpen < 5 {
		maxOpen = 5
	}

	if maxIdle < 5 {
		maxIdle = 5
	}

	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)

	return db, nil
}

func GetDb() *sql.DB {
	if db == nil {
		panic("config is not initial, should call InitMysql function")
	}
	return db
}
