package main

import (
	"github.com/Siddheshk02/jwt-auth-api/database"
	"github.com/Siddheshk02/jwt-auth-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.DBconn()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
