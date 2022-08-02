package main

import (
	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.UserRoutes(app)
	routes.StampRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(":3000")

}
