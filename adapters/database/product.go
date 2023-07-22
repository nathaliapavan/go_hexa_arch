package database

import (
	"database/sql"
	"github.com/nathaliapavan/go_hexa_arch/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDatabase struct {
	db *sql.DB
}

func NewProductDatabase(db *sql.DB) *ProductDatabase {
	return &ProductDatabase{db: db}
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

func (p *ProductDatabase) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select count(*) from products where id=?", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDatabase) create(product application.ProductInterface) (application.ProductInterface, error) {
	statement, err := p.db.Prepare("insert into products (id, name, price, status) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	err = statement.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDatabase) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update products set name=?, price=?, status=? where id=?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}