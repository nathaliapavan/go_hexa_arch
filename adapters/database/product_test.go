package database_test

import (
	"database/sql"
	"github.com/nathaliapavan/go_hexa_arch/adapters/database"
	"github.com/nathaliapavan/go_hexa_arch/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)


var Database *sql.DB

func setUp() {
	Database, _ = sql.Open("sqlite3", ":memory:")
	createTable(Database)
	createProduct(Database)
}

func createTable(database *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`
	statement, err := database.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func createProduct(database *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "Product Test", 0, "disabled");`
	statement, err := database.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func TestProductDatabase_Get(t *testing.T) {
	setUp()
	defer Database.Close()
	productDatabase := database.NewProductDatabase(Database)
	product, err := productDatabase.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDatabase_Save(t *testing.T) {
	setUp()
	defer Database.Close()
	productDatabase := database.NewProductDatabase(Database)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 50
	productResult, err := productDatabase.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDatabase.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
