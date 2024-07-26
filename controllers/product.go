package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	//
}

func (p Product) Index(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"": "",
	})
}

func (p Product) Show(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"": "",
	})
}

func (p Product) Store(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"": "",
	})
}

func (p Product) Update(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"": "",
	})
}

func (p Product) Delete(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"": "",
	})
}
