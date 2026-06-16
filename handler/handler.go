package handler

import (
	"walesp3982/golang-post-api/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func New(service services.Service) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	RoutesAuth(app, service.Auth)
	return app
}
