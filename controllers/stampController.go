package controllers

import (
	"net/http"

	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/models"
	"github.com/gofiber/fiber/v2"
)

func GetStamps(c *fiber.Ctx) error {
	stamps := []models.Stamp{}

	err := database.Database.Db.Find(&stamps).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get stamps"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "stamps fetched successfully", "data": stamps})
	return nil
}

func GetStamp(c *fiber.Ctx) error {
	stamp := models.Stamp{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	err := database.Database.Db.Where("id = ?", id).First(&stamp).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get the stamp from database"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Stamp has been fetched successfully",
		"data": stamp,
	})
	return nil
}

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
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Stamp has been added"})
	return nil
}

func DeleteStamp(c *fiber.Ctx) error {
	stamp := models.Stamp{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	err := database.Database.Db.Where("id = ?", id).Delete(&stamp).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not delete a stamp"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Stamp has been deleted"})
	return nil

}

func UpdateStamp(c *fiber.Ctx) error {
	stamp := models.Stamp{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	if err := c.BodyParser(&stamp); err != nil {
		c.Status(http.StatusServiceUnavailable).JSON(&fiber.Map{"message": err.Error()})
		return err
	}

	err := database.Database.Db.Where("id = ?", id).Updates(&stamp).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not update a stamp"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Stamp has been updated successfully"})
	return nil
}
