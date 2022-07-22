package routes

import (
	"github.com/Buniekbua/postage-stamps-managment/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUser)
	app.Post("/users", controllers.CreateUser)
	// app.Post("/users/login", controllers.Login)
	app.Delete("/users/:id", controllers.DeleteUser)
}
