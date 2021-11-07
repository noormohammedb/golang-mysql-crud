package controllers

import (
	"database/sql"
	"fmt"
)

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", DbCredentials())
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected")
	return db, nil
}
