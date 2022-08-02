package routes

import (
	"github.com/Buniekbua/postage-stamps-managment/controllers"
	"github.com/Buniekbua/postage-stamps-managment/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", middleware.Authenticate, controllers.GetUsers)
	app.Get("/users/:id", middleware.Authenticate, controllers.GetUser)
	app.Delete("/users/:id", middleware.Authenticate, controllers.DeleteUser)
	//app.Use(middleware.Authenticate)
}
