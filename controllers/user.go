package controllers

import "github.com/gofiber/fiber/v2"

type User struct {
	//
}

func (p User) Index(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (p User) Show(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (p User) Store(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (p User) Update(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (p User) Delete(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}
