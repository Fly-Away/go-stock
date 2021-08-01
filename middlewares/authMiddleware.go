package middlewares

import (
	"github.com/Fly-Away/go-stock/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		return util.ErrorResponse(c, err.Error(), "9999")
	}

	return c.Next()
}