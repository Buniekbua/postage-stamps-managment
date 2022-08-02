package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "error creating user"})
		return err
	}
	//Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Nickname: data["nickname"],
		Password: password,
		Email:    data["email"],
		Phone:    data["phone"],
	}

	//Check if email or phone exist in database
	emailExists := database.Database.Db.Where("email", data["email"]).First(&user)
	phoneExists := database.Database.Db.Where("phone", data["phone"]).First(&user)
	if emailExists.RowsAffected > 0 || phoneExists.RowsAffected > 0 {
		fmt.Println(emailExists.RowsAffected)
		fmt.Println(phoneExists.RowsAffected)
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "this email or phone already exists",
		})
	}

	err := database.Database.Db.Create(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "user has been created successfully",
		"data": user,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	user := models.User{}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	err := database.Database.Db.Where("email = ?", data["email"]).First(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not find user"})
		return err
	}
	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not find user"})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "incorrect password"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")
	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "incorrect password"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "success"})

}
