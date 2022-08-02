package routes

import (
	"github.com/Buniekbua/postage-stamps-managment/controllers"
	"github.com/Buniekbua/postage-stamps-managment/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.Login)
	app.Use("/validate", middleware.Authenticate)
}
