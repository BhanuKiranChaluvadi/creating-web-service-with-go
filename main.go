package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/inventoryservice/database"
	"github.com/pluralsight/inventoryservice/product"
	"github.com/pluralsight/inventoryservice/receipt"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	receipt.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5123", nil))
}
