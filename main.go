package main

import (
	"fmt"
	"gomerce/lib"
	"gomerce/models"
	"gomerce/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	if err := lib.InitENV(); err != nil {
		fmt.Println(err)
		return
	}

	if err := models.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Mount("/api", routers.New())

	app.Listen(":3000")
}
