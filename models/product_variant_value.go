package models

type ProductVariantValue struct {
	Id             int    `db:"id"`
	Value          string `db:"value"`
	Variant_key_id int    `db:"variant_key_id"`
}
