package controllers

import (
	"github.com/Fly-Away/go-stock/models"
	"github.com/Fly-Away/go-stock/models/dto/request/authDto"
	"github.com/Fly-Away/go-stock/services"
	"github.com/Fly-Away/go-stock/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

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

	user, err, errorCode, generatedToken := services.Login(userLoginRequest)
	if err != nil {
		return util.ErrorResponse(c, err.Error(), errorCode)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    *generatedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "Login success",
		Data: fiber.Map{
			"user": user,
		},
	})
}

// AuthenticateUser godoc
// @Summary get User authentication, this will get the jwt cookie
// @Description all field should filled
// @Accept  json
// @Produce  json
// @Tags Auth
// @Success 200 {object} models.SuccessResponseDto{}
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/user [get]
func AuthenticateUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	user, err, errorCode := services.AuthenticateUser(cookie)
	if err != nil {
		return util.UnAuthorizedResponse(c, err.Error()+". Not authenticated ", errorCode)
	}

	return c.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "User is valid",
		Data:    user,
	})
}
