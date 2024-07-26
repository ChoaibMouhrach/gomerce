package models

type ShippingOption struct {
	Id    int     `db:"id"`
	Name  string  `db:"name"`
	Price float32 `db:"price"`
}
