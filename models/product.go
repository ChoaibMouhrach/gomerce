package models

import "gomerce/lib"

type Product struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       uint   `db:"price"`
	Stock       uint   `db:"stock"`
}

type ProductCreate struct {
	Id          *string
	Name        string
	Description string
	Price       uint
	Stock       *uint
}

func AllProducts() ([]Product, error) {
	var products []Product

	err := lib.DB.Select(&products, "SELECT * FROM products")

	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func FirstProduct(id uint64) (*Product, error) {
	var product *Product

	err := lib.DB.Get(&product, "SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	return product, nil
}

func CreateProduct(input ProductCreate) error {
	_, err := lib.DB.Exec("INSERT INTO products(id, name, description, price, stock) VALUES(?, ?, ?, ?, ?)", input.Id, input.Name, input.Description, input.Price, input.Stock)
	return err
}
