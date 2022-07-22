package routes

import (
	"github.com/Buniekbua/postage-stamps-managment/controllers"
	"github.com/gofiber/fiber/v2"
)

func StampRoutes(app *fiber.App) {
	app.Get("/stamps", controllers.GetStamps)
	app.Get("/stamps/:id", controllers.GetStamp)
	app.Post("/stamps", controllers.CreateStamp)
	app.Delete("/stamps/:id", controllers.DeleteStamp)
	app.Patch("/stamps/:id", controllers.UpdateStamp)
}
