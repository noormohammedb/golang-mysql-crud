package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/noormohammedb/golang-mysql-crud/Controllers"
)

func Routers() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", ServerHome).Methods("get")
	muxRouter.HandleFunc("/get", GetAllData).Methods("get")
	return muxRouter
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Home")
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := controllers.GetDataFromDb()
	json.NewEncoder(w).Encode(data)
}
