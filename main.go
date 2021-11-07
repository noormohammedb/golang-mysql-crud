package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("golang mysql crud app")

	// insertData(db)
	// dataFromDb, _ := showPeaples(db)
	// fmt.Println(dataFromDb)

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

// func showPeaples(db *sql.DB) {
// }
// for dbResults {
// 	var row Person
// 	err := dbResults.Scan(&row)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(row)
// }

// }
