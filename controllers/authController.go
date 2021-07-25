package controllers

import (
	"github.com/Fly-Away/go-stock/models"
	"github.com/Fly-Away/go-stock/models/dto/request/authDto"
	"github.com/Fly-Away/go-stock/services"
	"github.com/Fly-Away/go-stock/util"
	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary User register, make sure email and password is unique
// @Description all field should filled
// @Accept  json
// @Produce  json
// @Tags Auth
// @Param item body authDto.UserRegisterRequest{} true "Post Hello yyy"
// @Success 200 {object} models.SuccessResponseDto{}
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/register [post]
func Register(c *fiber.Ctx) error {
	user := authDto.UserRegisterRequest{}

	if err := c.BodyParser(&user); err != nil {
		return util.ErrorResponse(c, err.Error(), 1001)
	}

	userResponse, err, errorCode := services.Register(user)

	if err != nil {
		return util.ErrorResponse(c, err.Error(), errorCode)
	}

	return c.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "Success registered",
		Data: fiber.Map{
			"user": userResponse,
		},
	})
}

// Login godoc
// @Summary User login, just login
// @Description all field should filled
// @Accept  json
// @Produce  json
// @Tags Auth
// @Param item body authDto.UserLoginRequest{} true "minimal 8 characters"
// @Success 200 {object} models.SuccessResponseDto{}
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/login [post]
func Login(c *fiber.Ctx) error {
	userLoginRequest := authDto.UserLoginRequest{}

	if err := c.BodyParser(&userLoginRequest); err != nil {
		return util.ErrorResponse(c, err.Error(), 1002)
	}

	user, err, errorCode := services.Login(userLoginRequest)
	if err != nil {
		return util.ErrorResponse(c, err.Error(), errorCode)
	}

	return c.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "Login success",
		Data: fiber.Map{
			"user": user,
		},
	})
}