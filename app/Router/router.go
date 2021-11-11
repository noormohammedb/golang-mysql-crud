package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/noormohammedb/golang-mysql-crud/app/Controllers"
	models "github.com/noormohammedb/golang-mysql-crud/app/Models"
)

func Routers() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", ServerHome).Methods("get")
	muxRouter.HandleFunc("/get", GetAllData).Methods("get")
	muxRouter.HandleFunc("/get/{id}", GetAData).Methods("get")
	muxRouter.HandleFunc("/add", AddPerson).Methods("post")
	muxRouter.HandleFunc("/seed", SeedData).Methods("get")
	muxRouter.HandleFunc("/edit/{id}", EditAPerson).Methods("put")
	muxRouter.HandleFunc("/delete/all", DeleteAllPerson).Methods("delete")
	muxRouter.HandleFunc("/delete/{id}", DeleteAPerson).Methods("delete")

	return muxRouter
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Home")
}

func SeedData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Seeding Data To Database")
	err := controllers.SeedDataToDb(controllers.GetData())
	if err != nil {
		json.NewEncoder(w).Encode("Seeding Error")
		return
	}
	json.NewEncoder(w).Encode("Seeding Success")
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all data")
	w.Header().Set("Content-Type", "application/json")
	data, _ := controllers.GetAllDataFromDb()
	json.NewEncoder(w).Encode(data)
}

func GetAData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all data")
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	data, _ := controllers.GetADataFromDb(parms["id"])
	json.NewEncoder(w).Encode(data)
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add a person")
	w.Header().Set("Content-Type", "application/json")
	var newPerson models.Person
	json.NewDecoder(r.Body).Decode(&newPerson)
	err := controllers.AddPersonData(newPerson)
	if err != nil {
		json.NewEncoder(w).Encode("error")
		return
	}
	json.NewEncoder(w).Encode("Success")
}

func EditAPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("edit person")
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	var editData models.Person
	decError := json.NewDecoder(r.Body).Decode(&editData)
	if parms["id"] == "" || decError != nil {
		json.NewEncoder(w).Encode("Request Error")
		return
	}
	rowCound, updateErr := controllers.EditPersonData(parms["id"], editData)
	if updateErr != nil {
		fmt.Println("update error")
		json.NewEncoder(w).Encode("Server Error")
		return
	}
	if rowCound > 0 {
		json.NewEncoder(w).Encode("Success")
		return
	}
	json.NewEncoder(w).Encode("No Change")
}

func DeleteAPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete person")
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	dbRowNos, err := controllers.DeleteAPerson(parms["id"])
	if err != nil {
		fmt.Println("err")
		json.NewEncoder(w).Encode("Server Error")
		return
	} else if dbRowNos > 0 {
		json.NewEncoder(w).Encode("Deleted")
		return
	}
	json.NewEncoder(w).Encode("No Change")
}

func DeleteAllPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete all person")
	w.Header().Set("Content-Type", "application/json")
	dbRows, err := controllers.DeleteAllPerson()
	fmt.Println(dbRows)
	if err != nil {
		fmt.Println("err")
		json.NewEncoder(w).Encode("Server Error")
		return
	} else if dbRows > 0 {
		json.NewEncoder(w).Encode("Deleted")
		return
	}
	json.NewEncoder(w).Encode("No Change")
}
