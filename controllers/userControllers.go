package controllers

import (
	"github.com/Fly-Away/go-stock/models"
	"github.com/Fly-Away/go-stock/services"
	"github.com/Fly-Away/go-stock/util"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
// @Summary get all user from database
// @Description no param needed
// @Accept  json
// @Produce  json
// @Tags user-module
// @Success 200 {object} models.SuccessResponseDto{}
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/user-module/view-all [get]
func GetAllUsers(c *fiber.Ctx) error {

	users, err := services.GetAllUsers()
	if err != nil {
		return util.ErrorResponse(c, err.Error(), 1101)
	}

	return c.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "Success",
		Data: fiber.Map{
			"users": users,
		},
	})
}