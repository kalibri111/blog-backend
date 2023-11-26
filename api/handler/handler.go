package handler

import (
	"backend/api"
	"backend/database/connection"
	"backend/database/service"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"time"
)

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Find the user by credentials
	user, err := service.GetUserByCreds(connection.DatabaseConnectionGlobal, loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(api.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(LoginResponse{
		Token: t,
	})
}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	favPhrase := claims["fav"].(string)
	return c.SendString("Welcome 👋" + email + " " + favPhrase)
}

func Logout(c *fiber.Ctx) error {
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"exp": time.Now(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(api.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(LoginResponse{
		Token: t,
	})
}

func HandleAllArticles(c *fiber.Ctx) error {
	articles, err := service.GetAllArticles(connection.DatabaseConnectionGlobal)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(articles)
}

func HandleAllPhotos(c *fiber.Ctx) error {
	photos, err := service.GetAllPhotos(connection.DatabaseConnectionGlobal)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(photos)
}

func HandleArticlesByHandler(c *fiber.Ctx) error {
	articles, err := service.GetArticleByHeader(connection.DatabaseConnectionGlobal, c.Params("header"))
	if err != nil {
		return err
	}
	return c.Status(200).JSON(articles)
}

func HandlePhotoByDate(c *fiber.Ctx) error {
	photos, err := service.GetPhotosByDate(connection.DatabaseConnectionGlobal, c.Params("date"))
	if err != nil {
		return err
	}
	return c.Status(200).JSON(photos)
}
