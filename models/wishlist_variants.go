package models

type WishlistVariant struct {
	Id        int `db:"id"`
	UserId    int `db:"user_id"`
	VariantId int `db:"variant_id"`
}
