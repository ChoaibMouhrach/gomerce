package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/resend/resend-go/v2"

	"gomerce/lib"
	"gomerce/models"
)

type Auth struct {
	//
}

type SignInSchema struct {
	Email string `json:"email" validate:"required,email"`
}

func (a Auth) SignIn(c *fiber.Ctx) error {
	input := new(SignInSchema)

	if err := lib.Validate(c, input); err != nil {
		return c.JSON(err)
	}

	user, err := models.UserFirstOrCreate(input.Email)

	if err != nil {
		return c.JSON(err.Error())
	}

	token, err := lib.GenerateToken(32)

	if err != nil {
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	client := resend.NewClient("")

	params := &resend.SendEmailRequest{
		From:    "auth@yerapos.com",
		To:      []string{user.Email},
		Subject: "Hello from Golang",
		Html:    fmt.Sprintf("<a href='%s'>Sign In</a>", token),
	}

	if _, err := client.Emails.Send(params); err != nil {
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	println(token)

	return c.JSON(input)
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
