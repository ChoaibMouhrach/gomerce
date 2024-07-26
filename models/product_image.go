package models

type ProductImage struct {
	Id        int    `db:"id"`
	Key       string `db:"key"`
	ProductId int    `db:"product_id"`
}
