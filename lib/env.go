package lib

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvStruct struct {
	// API
	API_URL string `validate:"required,url"`

	// APP
	APP_MAGIC_LINK string `validate:"required,url"`

	// RESEND
	RESEND_DOMAIN string `validate:"required"`
	RESEND_KEY    string `validate:"required"`
}

var Env EnvStruct

func InitENV() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	env := EnvStruct{
		// API
		API_URL: os.Getenv("API_URL"),

		// APP
		APP_MAGIC_LINK: os.Getenv("APP_MAGIC_LINK"),

		// RESEND
		RESEND_DOMAIN: os.Getenv("RESEND_DOMAIN"),
		RESEND_KEY:    os.Getenv("RESEND_KEY"),
	}

	if err := Validate(env); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Env = env

	return nil
}
