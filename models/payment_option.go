package models

type PaymentOption struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
