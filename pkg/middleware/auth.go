package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "auth-token" {
		return c.Next()
	} else {
		return c.JSON("Wrong auth token")
	}
}
