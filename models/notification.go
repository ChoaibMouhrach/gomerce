package models

type Notification struct {
	Id          int     `db:"id "`
	Title       string  `db:"title "`
	Description *string `db:"description"`
	Read        bool    `db:"read "`
	UserId      *int    `db:"user_id"`
	TypeId      int     `db:"type_id"`
}
