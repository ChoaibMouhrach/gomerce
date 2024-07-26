package models

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
