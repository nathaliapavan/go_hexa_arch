package database_test

import (
	"database/sql"
	"log"
)


var Database *sql.DB

func setUp() {
	Database, _ = sql.Open("sqlite3", ":momory:")
	createTable(Database)
	createProduct(Database)
}

func createTable(database *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string,
	);`
	statement, err := database.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func createProduct(database *sql.DB) {
	insert := `INSERT INTO products VALUES ("abd", "Product Test", 0, "disabled");`
	statement, err := database.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}