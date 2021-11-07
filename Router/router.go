package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/noormohammedb/golang-mysql-crud/Controllers"
	models "github.com/noormohammedb/golang-mysql-crud/Models"
)

func Routers() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", ServerHome).Methods("get")
	muxRouter.HandleFunc("/get", GetAllData).Methods("get")
	muxRouter.HandleFunc("/add", AddPerson).Methods("post")
	muxRouter.HandleFunc("/seed", SeedData).Methods("get")

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
	data, _ := controllers.GetDataFromDb()
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
