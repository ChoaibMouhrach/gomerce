package models

type Product struct {
	Id          int     `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`
	Thumbnail   *string `db:"thumbnail"`
	UnitId      int     `db:"unit_id"`
}
