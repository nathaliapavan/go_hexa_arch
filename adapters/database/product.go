package database

import (
	"database/sql"
	"github.com/nathaliapavan/go_hexa_arch/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDatabase struct {
	db *sql.DB
}

func (p *ProductDatabase) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	statement, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}