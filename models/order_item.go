package models

type OrderItem struct {
	Id        int     `db:"id"`
	Tax       float32 `db:"tax"`
	Price     float32 `db:"price"`
	Quantity  int     `db:"quantity"`
	OrderId   int     `db:"order_id"`
	VariantId int     `db:"variant_id"`
}
