package handler

import (
	"strings"
	"time"
	"walesp3982/golang-post-api/services"

	"github.com/gofiber/fiber/v3"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleRegister(service services.AuthService) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		req := new(RegisterRequest)
		err := c.Bind().Body(req)
		if err != nil {
			message := map[string]string{"message": "invalid data"}
			return c.Status(404).JSON(message)
		}
		user, err := service.Register(c.Context(), req.Name, req.Email, req.Password)

		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		return c.Status(200).JSON(user)
	}
}

type Credencials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handlerLogin(service services.AuthService) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		req := new(Credencials)
		err := c.Bind().Body(req)

		if err != nil {
			msg := map[string]string{"message": "invalid body"}
			return c.Status(400).JSON(msg)
		}

		refreshToken, err := service.Login(c.Context(), req.Email, req.Password)

		if err != nil {
			return c.Status(401).SendString(err.Error())
		}

		cookie := new(fiber.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = refreshToken
		cookie.Expires = time.Now().Add(time.Hour * 24 * 7)
		cookie.Secure = false
		cookie.HTTPOnly = true
		cookie.Path = "/auth"

		c.Cookie(cookie)

		return c.Status(202).SendString("")
	}
}

func handleAccessToken(service services.AuthService) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		// Get refresh token between cookies
		refreshToken := c.Cookies("refresh_token")
		if strings.Trim(refreshToken, " ") == "" {
			msg := map[string]string{"message": "Cookie not found error"}
			return c.Status(404).JSON(msg)
		}
		accessToken, err := service.GetAccessToken(c.Context(), refreshToken)

		if err != nil {
			msg := map[string]string{"message": err.Error()}
			return c.Status(400).JSON(msg)
		}

		return c.Status(200).JSON(map[string]string{"access_token": accessToken})
	}
}

func handleLogout(service services.AuthService) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		refreshToken := c.Cookies("refresh_token")
		if strings.Trim(refreshToken, " ") == "" {
			msg := map[string]string{"message": "Cookie not found error"}
			return c.Status(404).JSON(msg)
		}
		err := service.Logout(c.Context(), refreshToken)
		if err != nil {
			message := map[string]string{"message": err.Error()}
			return c.Status(400).JSON(message)
		}

		cookie := new(fiber.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = ""
		cookie.Expires = time.Now().Add(-time.Hour)
		cookie.Secure = false
		cookie.HTTPOnly = true
		cookie.Path = "/auth"

		c.Cookie(cookie)

		return c.Status(202).End()
	}
}

func RoutesAuth(app *fiber.App, service services.AuthService) {
	group := app.Group("/auth")

	group.Post("/register", handleRegister(service))
	group.Post("/login", handlerLogin(service))
	group.Post("/token", handleAccessToken(service))
	group.Post("/logout", handleLogout(service))
}
