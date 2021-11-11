package controllers

import (
	"database/sql"
	"fmt"

	config "github.com/noormohammedb/golang-mysql-crud/config"
)

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DbCredentials())
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected")
	return db, nil
}
