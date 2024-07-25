package routers

import (
	"gomerce/controllers"
	"gomerce/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter() *fiber.App {
	app := fiber.New()

	app.Use(middlewares.Auth)

	// AUTH
	auth := controllers.Auth{}
	app.Post("/sign-out", auth.SignOut)
	app.Get("/profile", auth.Profile)

	// USER
	user := controllers.User{}
	app.Get("/users", user.Index)
	app.Get("/users/:id", user.Show)
	app.Post("/users", user.Store)
	app.Patch("/users/:id", user.Update)
	app.Delete("/users/:id", user.Delete)

	// PRODUCT
	product := controllers.Product{}
	app.Get("/products", product.Index)
	app.Get("/products/:id", product.Show)
	app.Post("/products", product.Store)
	app.Patch("/products/:id", product.Update)
	app.Delete("/products/:id", product.Delete)

	return app
}
