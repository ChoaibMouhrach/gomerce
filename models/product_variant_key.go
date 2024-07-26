package models

type ProductVariantKey struct {
	Id  int    `db:"id"`
	Key string `db:"key"`
}
