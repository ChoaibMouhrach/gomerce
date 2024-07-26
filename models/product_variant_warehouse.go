package models

type ProductVariantWarehouse struct {
	Id           int `db:"id"`
	Stock        int `db:"stock"`
	Warehouse_id int `db:"warehouse_id"`
	Variant_id   int `db:"variant_id"`
}
