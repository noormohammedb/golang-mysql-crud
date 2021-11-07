package controllers

import (
	"fmt"
	"log"

	models "github.com/noormohammedb/golang-mysql-crud/Models"
)

func GetDataFromDb() ([]models.Person, error) {
	db, _ := GetDb()
	defer db.Close()
	dbQuery := "SELECT * FROM person"
	selectQuery, err := db.Query(dbQuery)
	if err != nil {
		panic(err)
	}
	var peaple []models.Person
	defer selectQuery.Close()
	for selectQuery.Next() {
		var iterPerson models.Person
		err := selectQuery.Scan(&iterPerson.Id, &iterPerson.Name, &iterPerson.Age, &iterPerson.Location)
		if err != nil {
			log.Fatal(err)
			return peaple, err
		}
		peaple = append(peaple, iterPerson)
	}
	return peaple, nil
}

func AddPersonData(person models.Person) error {
	db, _ := GetDb()
	defer db.Close()
	fmt.Println(person)
	dbQuery := "INSERT INTO person(name, age, location) VALUES(?, ?, ?)"
	insertQuery, err := db.Prepare(dbQuery)
	if err != nil {
		fmt.Println("insert query prepare error")
		return err
	}
	_, errIns := insertQuery.Exec(person.Name, person.Age, person.Location)
	if errIns != nil {
		return errIns
	}
	defer insertQuery.Close()
	return nil
}

func SeedDataToDb(person []models.Person) error {
	db, _ := GetDb()
	defer db.Close()
	dataToDb := GetData()
	dbInsertQuery := "INSERT INTO person(name, age, location) VALUES(?, ?, ?)"
	for _, data := range dataToDb {
		insertQuery, err := db.Prepare(dbInsertQuery)
		if err != nil {
			fmt.Println("insert query prepare error")
			return err
		}
		defer insertQuery.Close()
		insertQuery.Exec(data.Name, data.Age, data.Location)
	}
	return nil
}
