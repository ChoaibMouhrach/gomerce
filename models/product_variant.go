package models

type ProductVariant struct {
	Id         int     `db:"id"`
	Price      float32 `db:"price"`
	Tax        float32 `db:"tax"`
	Product_id int     `db:"product_id"`
}
