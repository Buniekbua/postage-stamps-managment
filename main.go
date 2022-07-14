package main

import (
	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	routes.UserRoutes(app)
	routes.StampRoutes(app)

	app.Listen(":3000")

}
