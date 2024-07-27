package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type MagicLink struct {
	Id      int       `db:"id"`
	Token   string    `db:"token"`
	UserId  int       `db:"user_id"`
	Expires time.Time `db:"expires"`
}

func (m MagicLink) Insert(tx ...*sqlx.Tx) error {
	var db sqlx.Execer = DB

	if len(tx) > 0 {
		db = tx[0]
	}

	_, err := db.Exec(
		`INSERT INTO magic_links(token, user_id, expires) VALUES($1, $2, $3)`,
		m.Token,
		m.UserId,
		time.Now().Add(60*time.Minute),
	)

	return err
}
