package controllers

import "github.com/gofiber/fiber/v2"

type Auth struct {
	//
}

func (a Auth) SignIn(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (a Auth) SignOut(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (a Auth) Profile(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}
