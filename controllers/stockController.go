package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// HelloWorld godoc
// @Summary test
// @Description testttt
// @Accept  json
// @Produce  json
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router / [get]
func HelloWorld(c *fiber.Ctx) error {
	// Create new Item and returns it
	return c.JSON(Item{
		Id: "1",
	})
}

// PostHelloWorld godoc
// @Summary This is post example
// @Description Just another test
// @Accept  json
// @Produce  json
// @Tags Item
// @Param item body Item true "Post Hello xx"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /this-is-our-new-api [post]
func PostHelloWorld(c *fiber.Ctx) error {
	// Create new Item and returns it
	return c.JSON(Item{
		Id: "1",
	})
}

// Item example
type Item struct {
	Id string `json:"id" example:"id something"`
}

type HTTPError struct {
	Status  string
	Message string
}