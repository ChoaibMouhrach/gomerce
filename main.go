package main

import (
	"fmt"
	"gomerce/models"
	"gomerce/routers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := models.InitDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	app := fiber.New()

	app.Mount("/api", routers.New())

	app.Listen(":3000")
}
