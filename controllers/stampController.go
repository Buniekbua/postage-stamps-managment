package controllers

import (
	"net/http"

	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/models"
	"github.com/gofiber/fiber/v2"
)

// func GetStamps(C *fiber.Ctx) error {

// }

// func GetStamp(C *fiber.Ctx) error {

// }

func CreateStamp(c *fiber.Ctx) error {
	stamp := new(models.Stamp)
	if err := c.BodyParser(stamp); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"error": "Request failed"})
		return err
	}
	err := database.Database.Db.Create(&stamp).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Could not create a stamp"})
		return err
	}
	//c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Stamp has been added"})
	return c.Status(http.StatusOK).JSON(&stamp)
}

// func DeleteStamp(C *fiber.Ctx) error {

// }

// func UpdateStamp(C *fiber.Ctx) error {

// }
