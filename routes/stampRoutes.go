package routes

import (
	"github.com/Buniekbua/postage-stamps-managment/controllers"
	"github.com/Buniekbua/postage-stamps-managment/middleware"
	"github.com/gofiber/fiber/v2"
)

func StampRoutes(app *fiber.App) {
	app.Get("/stamps", controllers.GetStamps)
	app.Get("/stamps/:id", controllers.GetStamp)
	app.Post("/stamps", middleware.Authenticate, controllers.CreateStamp)
	app.Delete("/stamps/:id", middleware.Authenticate, controllers.DeleteStamp)
	app.Patch("/stamps/:id", middleware.Authenticate, controllers.UpdateStamp)
}
