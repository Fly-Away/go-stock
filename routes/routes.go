package routes

import (
	"github.com/Fly-Away/go-stock/controllers"
	_ "github.com/Fly-Away/go-stock/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Add endpoint to serve swagger documentation
	app.Get("/swagger/*", swagger.Handler) // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
	}))

	mainApi := app.Group("/api")

	// Auth Endpoint
	v1 := mainApi.Group("/v1")
	v1.Post("/register", controllers.Register)
	v1.Post("/login", controllers.Login)
	v1.Get("/user", controllers.AuthenticateUser)
	// End of Auth Endpoint

	app.Get("/", controllers.HelloWorld)
	app.Post("/this-is-our-new-api", controllers.PostHelloWorld)
}