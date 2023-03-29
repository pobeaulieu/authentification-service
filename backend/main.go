package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // Authentification avec Cookie HTTP seulement.
	}))

	routes.Setup(app)

	app.Listen(":8000")

}