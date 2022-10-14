package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"url-shortner-service/entity"
	"url-shortner-service/routes"
)

func setupRoutes(app *fiber.App) {
	// API Routes
	app.Get("/api/url", routes.GetAll)
	app.Get("/api/url/:id", routes.Get)
	app.Post("/api/url/", routes.Create)
	// There shouldn't be an update API!
	app.Delete("/api/url/:id", routes.Delete)
}

func main() {
	entity.Setup()

	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")

}
