package middleware

import (
	"walesp3982/golang-post-api/services"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const NamespaceBookId = "postId"

func LoadBookId(service services.PostService, bookParameter string) fiber.Handler {
	return func(c fiber.Ctx) error {
		postStr := c.Params(bookParameter)
		postId, err := uuid.Parse(postStr)
		if err != nil {
			return c.Status(400).SendString("Post unknown")
		}
		fiber.Locals(c, NamespaceBookId, postId)
		return c.Next()
	}
}

func ContextPostId(c fiber.Ctx) uuid.UUID {
	return fiber.Locals[uuid.UUID](c, NamespaceBookId)
}

func HaveWriterPermission(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := ContextUserId(c)
		postId := ContextUserId(c)
		if !service.IsEditable(c, postId, userId) {
			return c.Status(fiber.StatusUnauthorized).SendString("Cannot have permission to edit this post")
		}

		return c.Next()
	}

}
