package lib

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	err := connect()

	if err != nil {
		return err
	}

	err = migrate()

	if err != nil {
		return err
	}

	return nil
}

var DB *sqlx.DB

func connect() error {
	database, err := sqlx.Connect("sqlite3", "./local.db")

	if err != nil {
		return err
	}

	DB = database

	return nil
}

func migrate() error {
	files, err := os.ReadDir("./migrations")

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, file := range files {
		content, err := os.ReadFile(fmt.Sprintf("./migrations/%s", file.Name()))

		if err != nil {
			return err
		}

		DB.MustExec(string(content))
	}

	return nil
}
