package handler

import (
	"backend/api"
	"backend/database/connection"
	"backend/database/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	jtoken "github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"
)

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(LoginRequest)
	loginRequest.Login = c.Query("Login")
	loginRequest.Password = c.Query("Password")
	// Find the user by credentials
	user, err := service.GetUserByCreds(connection.DatabaseConnectionGlobal, loginRequest.Login, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Login,
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
	return c.Status(200).JSON(LoginResponse{
		Token: t,
	})
}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	userDTO, err := service.GetUserByLogin(connection.DatabaseConnectionGlobal, email)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(userDTO)
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

func Register(c *fiber.Ctx) error {
	login := c.Query("Login")
	password := c.Query("Password")
	firstName := c.Query("FirstName", "aboba")
	lastName := c.Query("LastName", "abobus")
	user, err := service.GetUserByLogin(connection.DatabaseConnectionGlobal, login)

	if err != nil || user.Login == login {
		return c.Status(400).SendString("Login exist")
	}

	insertErr := service.AddUser(connection.DatabaseConnectionGlobal, login, password, firstName, lastName)

	if insertErr != nil {
		return c.Status(400).SendString("User insertion failure")
	}

	return c.Status(200).SendString("Registration OK")
}

func HandleAllArticles(c *fiber.Ctx) error {
	articles, err := service.GetAllArticles(connection.DatabaseConnectionGlobal)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(articles)
}

func HandleDownloadArticleFile(c *fiber.Ctx) error {
	articleId, errConv := strconv.Atoi(c.Query("id"))
	if errConv != nil {
		return errConv
	}

	article, err := service.GetArticleById(connection.DatabaseConnectionGlobal, articleId)

	if err != nil {
		return err
	}

	articlePath := article.Content

	errUpload := filesystem.SendFile(c, http.Dir("/Users/ivanleskin/Desktop/MIPT/fulstack/backend/"), articlePath)
	if errUpload != nil {
		// Handle the error, e.g., return a 404 Not Found response
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	return nil
}

func HandleDownloadPhotoFile(c *fiber.Ctx) error {
	photoId, errConv := strconv.Atoi(c.Query("id"))
	if errConv != nil {
		return errConv
	}

	photo, err := service.GetPhotoById(connection.DatabaseConnectionGlobal, photoId)

	if err != nil {
		return err
	}

	photoPath := photo.Photo

	errUpload := filesystem.SendFile(c, http.Dir("/Users/ivanleskin/Desktop/MIPT/fulstack/backend/"), photoPath)
	if errUpload != nil {
		// Handle the error, e.g., return a 404 Not Found response
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	return nil

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
