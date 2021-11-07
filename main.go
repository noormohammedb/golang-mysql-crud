package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	router "github.com/noormohammedb/golang-mysql-crud/Router"
)

func main() {
	fmt.Println("golang mysql CRUD app")

	log.Fatal(http.ListenAndServe(":8080", router.Routers()))
}
