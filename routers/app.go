package routers

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	app.Mount("/", AuthRouter())
	app.Mount("/", GuestRouter())

	return app
}
