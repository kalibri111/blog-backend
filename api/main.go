package api

import (
	"backend/api/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/article/all/", func(ctx *fiber.Ctx) error {
		return handler.HandleAllArticles(ctx)
	})

	app.Get("/article/filter/:header", func(ctx *fiber.Ctx) error {
		return handler.HandleArticlesByHandler(ctx)
	})

	app.Get("/photo/all/", func(ctx *fiber.Ctx) error {
		return handler.HandleAllPhotos(ctx)
	})

	app.Get("/photo/filter/:date", func(ctx *fiber.Ctx) error {
		return handler.HandlePhotoByDate(ctx)
	})

	app.Get("/user/login/", func(ctx *fiber.Ctx) error {
		return handler.Login(ctx)
	})

	app.Get("/user/protected/", func(ctx *fiber.Ctx) error {
		return handler.Protected(ctx)
	})

	app.Get("/user/logout/", func(ctx *fiber.Ctx) error {
		return handler.Logout(ctx)
	})

}
