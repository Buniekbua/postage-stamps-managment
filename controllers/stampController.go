package controllers

import (
	"fmt"
	"net/http"

	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/models"
	"github.com/gofiber/fiber/v2"
)

func GetStamps(c *fiber.Ctx) error {
	stamps := []models.Stamp{}

	err := database.Database.Db.Find(&stamps).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get stamps"})
		return err
	}
	responseStamps := []models.Stamp{}

	for _, stamp := range stamps {
		var user models.User
		database.Database.Db.Find(&user, "id = ?", stamp.UserRefer)
		stamp.User = user
		responseStamps = append(responseStamps, stamp)
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "stamps fetched successfully",
		"data":    responseStamps,
	})
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
	var user models.User
	database.Database.Db.First(&user, stamp.UserRefer)
	stamp.User = user
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "stamp has been fetched successfully",
		"data": stamp,
	})
	return nil
}

func CreateStamp(c *fiber.Ctx) error {
	stamp := models.Stamp{}
	if err := c.BodyParser(&stamp); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"error": "request failed"})
		return err
	}

	user := models.User{}
	id := c.Locals("id")
	fmt.Println(id)
	err2 := database.Database.Db.Where("id = ?", id).First(&user).Error
	if err2 != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not find a user"})
		return err2
	}

	stampModel := models.Stamp{
		Name:        stamp.Name,
		Description: stamp.Description,
		Price:       stamp.Price,
		Image:       stamp.Image,
		Quantity:    stamp.Quantity,
		User:        user,
	}

	err := database.Database.Db.Create(&stampModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "could not create a stamp"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "stamp has been added",
		"data":    stampModel,
	})
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
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete a stamp"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "stamp has been deleted"})
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
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not update a stamp"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "stamp has been updated successfully"})
	return nil
}
