package models

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
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

	err = seed()

	if err != nil {
		return err
	}

	return nil
}

var DB *sqlx.DB

func connect() error {
	dsn := "user=root password=3636 host=0.0.0.0 port=5432 dbname=gomerce sslmode=disable"
	database, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return err
	}

	DB = database

	return nil
}

func migrate() error {
	files, err := os.ReadDir("./migrations")

	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := os.ReadFile(fmt.Sprintf("./migrations/%s", file.Name()))

		if err != nil {
			return err
		}

		_, err = DB.Exec(string(content))

		if err != nil {
			return fmt.Errorf("- %s: %s", file.Name(), err.Error())
		}
	}

	return nil
}

func seedRoles() error {
	var roles []Role

	if err := DB.Select(&roles, "SELECT * FROM roles"); err != nil {
		return err
	}

	if len(roles) != 0 {
		return nil
	}

	if _, err := DB.Exec("INSERT INTO roles (name) VALUES ($1), ($2)", "member", "admin"); err != nil {
		return err
	}

	return nil
}

func seed() error {
	if err := seedRoles(); err != nil {
		return fmt.Errorf("- seed: %s", err.Error())
	}

	return nil
}
