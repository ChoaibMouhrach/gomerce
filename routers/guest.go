package routers

import (
	"gomerce/controllers"
	"gomerce/middlewares"

	"github.com/gofiber/fiber/v2"
)

func GuestRouter() *fiber.App {
	app := fiber.New()

	app.Use(middlewares.Guest)

	// AUTH
	auth := controllers.Auth{}
	app.Post("/sign-in", auth.SignIn)
	app.Post("/auth", auth.Authenticate)

	return app
}
