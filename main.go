package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("golang mysql crud app")
	db, err := sql.Open("mysql", DbCredentials())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// insertData(db)
	dataFromDb, _ := showPeaples(db)
	fmt.Println(dataFromDb)

	jsonData, _ := json.MarshalIndent(dataFromDb, "", " ")

	fmt.Println(string(jsonData))

}

// func insertData(db *sql.DB) error {
// 	peaple := GetData()
// 	fmt.Println(peaple)
// 	for _, person := range peaple {
// 		dbQuery := "INSERT INTO person(name, age, location) VALUES(?, ?, ?)"
// 		insertQuery, err := db.Prepare(dbQuery)
// 		if err != nil {
// 			fmt.Println("insert query prepare error")
// 			return err
// 		}
// 		defer insertQuery.Close()

// 		dbResponse, err := insertQuery.Exec(person.Name, person.Age, person.Location)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(dbResponse)
// 	}
// 	return nil

// }

func showPeaples(db *sql.DB) ([]Person, error) {
	dbQuery := "SELECT * FROM person"
	selectQuery, err := db.Query(dbQuery)
	if err != nil {
		panic(err)
	}
	var peaple []Person
	defer selectQuery.Close()
	for selectQuery.Next() {
		var iterPerson Person

		err := selectQuery.Scan(&iterPerson.Id, &iterPerson.Name, &iterPerson.Age, &iterPerson.Location)
		if err != nil {
			log.Fatal(err)
			return peaple, err
		}

		peaple = append(peaple, iterPerson)

	}
	// for dbResults {
	// 	var row Person
	// 	err := dbResults.Scan(&row)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(row)
	// }

	return peaple, nil
}
