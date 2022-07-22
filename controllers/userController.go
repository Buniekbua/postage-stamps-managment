package controllers

import (
	"net/http"

	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	err := database.Database.Db.Find(&users).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get users"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Users fetched successfully",
		"data": users,
	})
	return nil
}

func GetUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	err := database.Database.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not find user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been fetched successfully",
		"data": user,
	})
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	err := database.Database.Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not delete a user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been deleted successfully"})
	return nil
}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Error creating user"})
		return err
	}
	err := database.Database.Db.Create(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not create user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been created successfully"})
	return nil
}

// func Login(c *fiber.Ctx) error {

// }
