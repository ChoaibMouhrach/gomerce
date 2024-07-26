package models

type Order struct {
	Id                int     `db:"id"`
	ShippingPrice     float32 `db:"shipping_price"`
	UserId            int     `db:"user_id"`
	StatusId          int     `db:"status_id"`
	PaymentOptionId   int     `db:"payment_option_id"`
	ShippingOptionId  int     `db:"shipping_option_id"`
	ShippingAddressId int     `db:"shipping_address_id"`
}
