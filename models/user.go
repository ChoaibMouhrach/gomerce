package models

type User struct {
	Id     int    `db:"Id"`
	Name   string `db:"name"`
	Email  string `db:"email"`
	RoleId int    `db:"role_id"`
}
