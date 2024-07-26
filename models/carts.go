package models

type Cart struct {
	Id     int  `db:"id"`
	UserId *int `db:"user_id"`
}
