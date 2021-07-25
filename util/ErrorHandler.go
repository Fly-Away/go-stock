package util

import (
	"github.com/Fly-Away/go-stock/models"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, msg string, errorCode interface{}) error {
	return c.Status(400).JSON(models.ErrorResponse{
		Success: false,
		Message: msg,
		ErrorCode: errorCode,
		Url: c.Route().Path,
	})
}