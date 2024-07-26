package models

type CartVariant struct {
	Id        int `db:"id"`
	Quantity  int `db:"quantity"`
	CartId    int `db:"cart_id"`
	VariantId int `db:"variant_id"`
}
