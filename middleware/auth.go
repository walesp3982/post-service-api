package middleware

import (
	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/google/uuid"
)

const NamespaceUserId string = "userId"

func JWTMiddleware(key string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(key)},
		Extractor:  extractors.FromAuthHeader("Bearer"),
	})
}

// Load a userId of jwt in context Locals
func LoadUserIdContext(c fiber.Ctx) error {
	payload := jwtware.FromContext(c)
	strId, err := payload.Claims.GetSubject()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// Save uuid in locals
	id, err := uuid.Parse(strId)
	if err != nil {
		return c.Status(400).SendString("Invalid uuid in token")
	}
	fiber.Locals(c, NamespaceUserId, id)
	return c.Next()
}

func ContextUserId(c fiber.Ctx) uuid.UUID {
	return fiber.Locals[uuid.UUID](c, NamespaceUserId)
}
