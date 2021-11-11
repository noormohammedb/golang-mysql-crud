package controllers

import (
	"fmt"
	"log"

	models "github.com/noormohammedb/golang-mysql-crud/app/Models"
)

func GetAllDataFromDb() ([]models.Person, error) {
	db, _ := GetDb()
	defer db.Close()
	dbQuery := "SELECT id, Name, Age, Location FROM person WHERE isDeleted=0"
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

func GetADataFromDb(dataId string) ([]models.Person, error) {
	db, _ := GetDb()
	defer db.Close()
	dbQuery := "SELECT id, Name, Age, Location FROM person WHERE isDeleted=0 AND id=?"
	selectQuery, err := db.Prepare(dbQuery)
	if err != nil {
		panic(err)
	}
	selectResult, _ := selectQuery.Query(dataId)
	var peaple []models.Person
	defer selectQuery.Close()
	for selectResult.Next() {
		var iterPerson models.Person
		err := selectResult.Scan(&iterPerson.Id, &iterPerson.Name, &iterPerson.Age, &iterPerson.Location)
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

func EditPersonData(id string, newData models.Person) (int64, error) {
	db, _ := GetDb()
	defer db.Close()
	dbUpdateQuery := "UPDATE person SET name=?, age=?, location=? WHERE id=?"
	dataUpdate, err := db.Prepare(dbUpdateQuery)
	if err != nil {
		return 0, err
	}
	defer dataUpdate.Close()
	dbRes, dbUpdateErr := dataUpdate.Exec(newData.Name, newData.Age, newData.Location, id)
	if dbUpdateErr != nil {
		log.Fatal(dbUpdateErr)
		return 0, dbUpdateErr
	}
	rows, _ := dbRes.RowsAffected()
	return rows, nil
}

func DeleteAPerson(perId string) (int64, error) {
	db, _ := GetDb()
	defer db.Close()
	dbUpdateQuery := "UPDATE person SET isDeleted=1 WHERE id=?"
	dbUpdate, err := db.Prepare(dbUpdateQuery)
	ErrorCheck(err)
	defer dbUpdate.Close()
	dbRes, dbErr := dbUpdate.Exec(perId)
	ErrorCheck(dbErr)
	rows, _ := dbRes.RowsAffected()
	return rows, nil
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
