package models

type ShippingAddress struct {
	Id       int     `db:"id"`
	Phone    string  `db:"phone"`
	Address1 string  `db:"address_1"`
	Address2 *string `db:"address_2"`
	Billing  string  `db:"billing"`
	UserId   int     `db:"user_id"`
}
