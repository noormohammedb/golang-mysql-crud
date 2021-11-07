package controllers

import (
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
