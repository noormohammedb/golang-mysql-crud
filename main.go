package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	router "github.com/noormohammedb/golang-mysql-crud/Router"
)

func main() {
	fmt.Println("golang mysql crud app")

	log.Fatal(http.ListenAndServe(":8080", router.Routers()))

	// insertData(db)

	// jsonData, _ := json.MarshalIndent(dataFromDb, "", " ")

	// fmt.Println(string(jsonData))

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
