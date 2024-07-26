package models

type Order struct {
	Id                  int `db:"id"`
	user_id             int `db:"user_id"`
	status_id           int `db:"status_id"`
	payment_option_id   int `db:"payment_option_id "`
	shipping_option_id  int `db:"shipping_option_id "`
	shipping_address_id int `db:"shipping_address_id"`
}
