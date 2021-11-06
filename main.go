package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("golang mysql crud app")
	db, err := sql.Open("mysql", DbCredentials())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertData(db)
	// showPeaples(db)

}
func insertData(db *sql.DB) error {
	peaple := GetData()
	fmt.Println(peaple)
	for _, person := range peaple {
		dbQuery := "INSERT INTO person(name, age, location) VALUES(?, ?, ?)"
		insertQuery, err := db.Prepare(dbQuery)
		if err != nil {
			fmt.Println("insert query prepare error")
			return err
		}
		defer insertQuery.Close()

		dbResponse, err := insertQuery.Exec(person.Name, person.Age, person.Location)
		if err != nil {
			return err
		}
		fmt.Println(dbResponse)
	}
	return nil

}

// func showPeaples(db *sql.DB) {
// 	dbQuery := "SELECT * FROM person"
// 	selectQuery, _ := db.Prepare(dbQuery)
// 	_, err := selectQuery.Exec()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var per []Person
// 	defer selectQuery.Close()
// 	selectQuery.QueryRow().Scan(&per)
// 	fmt.Printf("%#v", selectQuery)
// 	fmt.Println(per)
// 	// for dbResults {
// 	// 	var row Person
// 	// 	err := dbResults.Scan(&row)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	fmt.Println(row)
// 	// }
// }
