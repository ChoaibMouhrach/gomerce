package controllers

import (
	"gomerce/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	//
}

func (p Product) Index(c *fiber.Ctx) error {
	products, err := models.AllProducts()

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(products)
}

func (p Product) Show(c *fiber.Ctx) error {
	rawId := c.Params("id")

	if len(rawId) == 0 {
		return c.Status(400).JSON(map[string]interface{}{
			"status": false,
		})
	}

	id, err := strconv.ParseUint(rawId, 10, 64)

	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status": false,
		})
	}

	product, err := models.FirstProduct(id)

	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status": false,
		})
	}

	return c.JSON(product)
}

func (p Product) Store(c *fiber.Ctx) error {
	models.CreateProduct(models.ProductCreate{
		Id:          nil,
		Name:        "I7 13700KF",
		Description: "This is an i7 Inel CPU",
		Price:       200,
		Stock:       nil,
	})

	return c.JSON(map[string]interface{}{
		"success": true,
	})
}

func (p Product) Update(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}

func (p Product) Delete(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"name": "choaib",
	})
}
