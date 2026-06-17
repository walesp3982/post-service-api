package handler

import (
	"walesp3982/golang-post-api/middleware"
	"walesp3982/golang-post-api/services"

	"github.com/gofiber/fiber/v3"
)

func RoutesUser(app *fiber.App, service services.UserService, jwtSecret string) {
	group := app.Group("/me")

	group.Use(
		middleware.JWTMiddleware(jwtSecret),
		middleware.LoadUserIdContext,
	)

	group.Get("/", handleCurrentUser(service))
	group.Post("/change-password", handleChangePasswordUser(service))
	group.Patch("/", handleUpdateUser(service))
	group.Delete("/", handleDeleteUser(service))

}

func handleCurrentUser(service services.UserService) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		id := middleware.ContextUserId(c)

		user := service.GetById(c, id)

		if user == nil {
			return c.Status(404).SendString("User not found")
		}

		return c.Status(200).JSON(user)
	}
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func handleChangePasswordUser(service services.UserService) fiber.Handler {
	return func(c fiber.Ctx) error {
		req := new(ChangePasswordRequest)

		id := middleware.ContextUserId(c)

		if err := c.Bind().Body(req); err != nil {
			return c.Status(400).SendString("Bad Request")
		}

		if err := service.ChangePassword(c, id, req.NewPassword, req.OldPassword); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

type UpdateUserRequest struct {
	Name string
}

func handleUpdateUser(service services.UserService) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get id in locals
		id := middleware.ContextUserId(c)

		req := new(UpdateUserRequest)
		if err := c.Bind().Body(req); err != nil {
			return c.Status(401).SendString("Bad request error")
		}

		if err := service.ChangeMetadata(c, id, req.Name); err != nil {
			return c.Status(400).SendString("Cannot change user error")
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func handleDeleteUser(service services.UserService) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := middleware.ContextUserId(c)

		if err := service.DeleteUser(c, id); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
