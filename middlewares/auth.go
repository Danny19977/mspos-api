package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("token")

	if _, err := utils.VerifyJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}