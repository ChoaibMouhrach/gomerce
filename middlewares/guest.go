package middlewares

import "github.com/gofiber/fiber/v2"

func Guest(c *fiber.Ctx) error {
	return c.Next()
}
