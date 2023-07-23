package main

import (
	"database/sql"
	db2 "github.com/nathaliapavan/go_hexa_arch/adapters/database"
	"github.com/nathaliapavan/go_hexa_arch/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDatabaseAdapter := db2.NewProductDatabase(db)
	productService := application.NewProductService(productDatabaseAdapter)
	product, _ := productService.Create("Product example", 50)
	productService.Enable(product)
}