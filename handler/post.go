package handler

import (
	"walesp3982/golang-post-api/middleware"
	"walesp3982/golang-post-api/services"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func RoutesPost(app *fiber.App, services services.Service, jwtSecret string) {
	userPosts := app.Group("/me/posts")
	userPosts.Use(
		middleware.JWTMiddleware(jwtSecret),
		middleware.LoadUserIdContext,
	)
	userPosts.Get("/", handleShowUserPosts(services.Post))
	userPosts.Post("/", handleCreatePost(services.Post))
	userPosts.Use(
		middleware.LoadBookId(services.Post, "postId"),
		middleware.HaveWriterPermission(services.Post),
	)
	userPosts.Get(
		"/:postId<guid>",
		handleGetPost(services.Post),
	)
	userPosts.Delete(
		"/:postId<guid>",
		handleDeletePost(services.Post),
	)
	userPosts.Put(
		"/:postId<guid>",
		handleUpdatePost(services.Post),
	)
	publicPosts := app.Group("/books")
	publicPosts.Get("/:slug", handleShowPublicBooks(services.Post))

}

func handleShowUserPosts(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := middleware.ContextUserId(c)

		posts := service.ShowPostsUser(c, userId)
		return c.Status(200).JSON(posts)
	}
}

func handleGetPost(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		bookId := middleware.ContextPostId(c)

		post := service.GetPost(c, bookId)
		if post == nil {
			return c.Status(fiber.StatusNotFound).SendString("Book not found")
		}

		return c.Status(200).JSON(post)

	}
}

func handleDeletePost(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		postId := middleware.ContextPostId(c)

		if err := service.Delete(c, postId); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func handleShowPublicBooks(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		pagination := GetPaginationQuery(c)
		books := service.ShowPosts(c, pagination.Offset, pagination.Limit)
		return c.JSON(books)
	}
}

type CreatePostRequest struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Slug    *string `json:"slug,omitempty"`
}

func handleCreatePost(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		userId := middleware.ContextUserId(c)
		req := new(CreatePostRequest)
		if err := c.Bind().Body(req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Json invalid")
		}
		post, err := service.Create(c, userId, req.Title, req.Content, req.Slug)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		if post == nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(post)
	}
}

type UpdatePostRequest struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Slug    *string   `json:"slug,omitempty"`
}

func handleUpdatePost(service services.PostService) fiber.Handler {
	return func(c fiber.Ctx) error {
		postId := middleware.ContextPostId(c)
		req := new(UpdatePostRequest)
		if err := c.Bind().Body(req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("JSON invalid error")
		}
		if postId != req.Id {
			return c.Status(fiber.StatusBadRequest).SendString("Id post invalid")
		}

		if err := service.UpdatePost(c, postId, req.Title, req.Content, req.Slug); err != nil {
			return c.Status(fiber.StatusBadGateway).SendString("Cannot update post")
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
