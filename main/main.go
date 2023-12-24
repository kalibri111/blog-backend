package main

import (
	"backend/api"
	"backend/database/connection"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	jwtware "github.com/gofiber/jwt/v3"
	"log"
	"net/http"

	"backend/api/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func main() {
	app := fiber.New()
	err := connection.NewSQLiteConnection()
	if err != nil {
		log.Fatal("Connection pizdec")
	}

	jwt := NewAuthMiddleware(api.Secret)

	app.Use(cors.New())

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("content_example"),
	}))

	app.Get("/article/all/", func(ctx *fiber.Ctx) error {
		return handler.HandleAllArticles(ctx)
	})

	app.Get("/article/download/", func(ctx *fiber.Ctx) error {
		return handler.HandleDownloadArticleFile(ctx)
	})

	app.Get("/article/filter/:header", func(ctx *fiber.Ctx) error {
		return handler.HandleArticlesByHandler(ctx)
	})

	app.Get("/photo/all/", func(ctx *fiber.Ctx) error {
		return handler.HandleAllPhotos(ctx)
	})

	app.Get("/photo/download/", func(ctx *fiber.Ctx) error {
		return handler.HandleDownloadPhotoFile(ctx)
	})

	app.Get("/photo/filter/:date", func(ctx *fiber.Ctx) error {
		return handler.HandlePhotoByDate(ctx)
	})

	app.Get("/user/login/", func(ctx *fiber.Ctx) error {
		return handler.Login(ctx)
	})

	app.Get("/user/protected/", jwt, func(ctx *fiber.Ctx) error {
		return handler.Protected(ctx)
	})

	app.Get("/user/logout/", func(ctx *fiber.Ctx) error {
		return handler.Logout(ctx)
	})

	app.Post("/user/register/", func(ctx *fiber.Ctx) error {
		return handler.Register(ctx)
	})

	log.Fatal(app.Listen(":4000"))

}
