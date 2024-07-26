package models

type NotificationType struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
