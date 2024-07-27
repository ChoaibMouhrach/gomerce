package controllers

import (
	"gomerce/lib"
	"gomerce/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	//
}

type signInSchema struct {
	Email string `json:"email" validate:"required,email"`
}

func (a Auth) SignIn(c *fiber.Ctx) error {
	input := new(signInSchema)
	if err := lib.ValidateRequest(c, input); err != nil {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: err,
		})
	}

	tx, err := models.DB.Beginx()
	if err != nil {
		return c.Status(500).JSON(lib.Response{
			Success: false,
			Message: "couldn't begin transaction",
		})
	}

	defer func() error {
		if err == nil {
			return nil
		}

		err = tx.Rollback()

		if err == nil {
			return nil
		}

		return c.Status(500).JSON(lib.Response{
			Success: false,
			Message: "Errror bitch",
		})
	}()

	var user models.User
	if err = tx.Get(&user, "SELECT * FROM users WHERE email=$1", user.Email); err != nil {

		var role models.Role
		if err := tx.Get(&role, "SELECT * FROM roles WHERE name=$1", models.MEMBER); err != nil {
			return c.Status(500).JSON(lib.Response{
				Success: false,
				Message: "role not found",
			})
		}

		if _, err = tx.Exec("INSERT INTO users (email, role_id) VALUES ($1, $2)", input.Email, role.Id); err != nil {
			return c.Status(500).JSON(lib.Response{
				Success: false,
				Message: "couldn't insert use",
			})
		}

		err = tx.Get(&user, "SELECT * FROM users WHERE email=$1", input.Email)
		if err != nil {
			return c.Status(500).JSON(lib.Response{
				Success: false,
				Message: "couldn't find user",
			})
		}

	}

	var token string
	if err := lib.GenerateToken(&token, 32); err != nil {
		return c.Status(500).JSON(lib.Response{
			Success: false,
			Message: "We couldn't generate your token",
		})
	}

	magicLink := models.MagicLink{
		Token:  token,
		UserId: user.Id,
	}

	if err := magicLink.Insert(tx); err != nil {
		return c.Status(500).JSON(lib.Response{
			Success: false,
			Message: "couldn't generate auth token",
		})
	}

	if err := lib.SendMagicLink(user.Email, token); err != nil {
		return c.Status(500).JSON(lib.Response{
			Success: false,
			Message: "couldn't send email",
		})
	}

	if err := tx.Commit(); err != nil {
		return c.Status(500).JSON(lib.Response{
			Success: false,
			Message: "couldn't commit transaction",
		})
	}

	return c.JSON(lib.Response{
		Success: true,
		Message: "Check your inbox for our auth email",
	})
}

type authenticationSchema struct {
	Token string `json:"token" validate:"min=16"`
}

// !todo: make sure to add tx tomorrow
func (a Auth) Authenticate(c *fiber.Ctx) error {
	input := new(authenticationSchema)
	if err := lib.ValidateRequest(c, input); err != nil {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: err,
		})
	}

	var token models.MagicLink
	if err := models.DB.Get(&token, "SELECT * FROM magic_links WHERE token=$1", input.Token); err != nil {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: "Token not found",
		})
	}

	if _, err := models.DB.Exec("DELETE FROM magic_links WHERE id=$1", token.Id); err != nil {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: "We couldn't delete this token",
		})
	}

	if time.Now().After(token.Expires) {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: "Token expired",
		})
	}

	var sessionId string
	if err := lib.GenerateToken(&sessionId, 32); err != nil {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: "We couldn't generate you a session",
		})
	}

	_, err := models.DB.Exec(`INSERT INTO sessions (
		session_id,
		expires
	 ) VALUES (
		$1,
		$2
	)`, sessionId, time.Now().Add(30*24*time.Hour))

	if err != nil {
		return c.Status(400).JSON(lib.Response{
			Success: false,
			Message: "We couldn't create you a new session",
		})
	}

	return c.JSON(lib.Response{
		Success: true,
		Message: map[string]string{
			"session_id": sessionId,
		},
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
