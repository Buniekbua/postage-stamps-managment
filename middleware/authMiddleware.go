package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Buniekbua/postage-stamps-managment/database"
	"github.com/Buniekbua/postage-stamps-managment/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {

	tokenString := c.Cookies("jwt")

	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"messege": "user unauthorized"})
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"messege": "user unauthorized"})
		}

		var user models.User
		errorDb := database.Database.Db.Where("id = ?", claims["iss"]).First(&user).Error
		if errorDb != nil {
			c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not find user"})
			return errorDb
		}

		if user.ID == 0 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"messege": "user unauthorized"})
		}

		c.Locals("id", user.ID)

		c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "user authorized",
			"data":    user,
		})

	} else {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"messege": "user unauthorized"})
	}

	c.Next()

	return nil
}
