package models

type PaymentOption struct {
	Id   int    `db:"id"`
	name string `db:"name"`
}
