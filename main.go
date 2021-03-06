package main

import (
	"log"
	"net/http"
	"simple-crud/database"
	"simple-crud/product"

	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
