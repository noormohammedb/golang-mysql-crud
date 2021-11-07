package controllers

import "database/sql"

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", DbCredentials())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db, nil
}
