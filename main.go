package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	// Sl       float64
	User string
	// Password string
}

func main() {
	fmt.Println("golang mysql crud app")
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(127.0.0.1)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	dbResults, err := db.Query("SELECT user FROM test_table")
	if err != nil {
		panic(err)
	}

	for dbResults.Next() {
		var row User

		err := dbResults.Scan(&row.User)
		if err != nil {
			panic(err)
		}

		fmt.Println(row)
	}

}
