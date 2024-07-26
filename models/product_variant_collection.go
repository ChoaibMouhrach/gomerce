package models

type ProductVariantCollection struct {
	Id             int `db:"id "`
	VariantId      int `db:"variant_id "`
	VariantKeyId   int `db:"variant_key_id"`
	VariantValueId int `db:"variant_value_id"`
}
