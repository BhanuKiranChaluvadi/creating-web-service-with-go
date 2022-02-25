package main

import (
	"log"
	"net/http"

	"github.com/pluralsight/inventoryservice/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5123", nil))
}
