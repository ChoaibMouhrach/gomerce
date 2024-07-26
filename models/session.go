package models

type Session struct {
	Id        int    `db:"id"`
	SessionId string `db:"session_id"`
	UserId    int    `db:"user_id"`
}
