package models

type User struct {
	Id     int     `db:"id"`
	Name   *string `db:"name"`
	Email  string  `db:"email"`
	Phone  *string `db:"phone"`
	RoleId int     `db:"role_id"`
}
