package services

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"walesp3982/golang-post-api/model"
	"walesp3982/golang-post-api/repository"

	"github.com/google/uuid"
)

type PostService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *PostService {
	return &PostService{
		repository: repository,
	}
}

func (s *PostService) GetPost(ctx context.Context, id uuid.UUID) *model.Post {
	post := s.repository.GetById(ctx, id)

	return post
}

func (s *PostService) ShowPosts(ctx context.Context, offset uint, limit uint) []model.Post {
	posts := s.repository.GetAllLastest(ctx, offset, limit)
	return posts
}

// Verify if the user have permission of editing
// (in other works what the user was created this post)
func (s *PostService) IsEditable(ctx context.Context, id uuid.UUID, userId uuid.UUID) bool {
	post := s.repository.GetById(ctx, id)
	if post == nil {
		return false
	}

	return post.UserId == userId
}

func (s *PostService) DeleteRepository(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}

func (s *PostService) FindBySlug(ctx context.Context, slug string) *model.Post {
	return s.repository.GetBySlug(ctx, slug)
}

func (s *PostService) ShowPostsUser(ctx context.Context, userId uuid.UUID) []model.Post {
	return s.repository.GetByUser(ctx, userId)
}

func (s *PostService) UpdatePost(ctx context.Context, post model.Post) error {
	return s.repository.Update(ctx, &post)
}

func GenerateSlug(title string) string {
	slug := strings.ToLower(title)
	// 2. Reemplazar todo lo que NO sea letras, números o espacios por vacío
	reg := regexp.MustCompile(`[^a-z0-9_-]`)
	slug = reg.ReplaceAllString(slug, "")

	// 3. Reemplazar espacios (y guiones bajos si prefieres) por un solo guion
	regSpace := regexp.MustCompile(`[\s_-]+`)
	slug = regSpace.ReplaceAllString(slug, "-")

	// 4. Quitar guiones sobrantes al principio y al final
	slug = strings.Trim(slug, "-")

	return slug
}
func (s *PostService) Create(ctx context.Context, userId uuid.UUID, title string, content string, slugOptional *string) (*model.Post, error) {
	var slug string
	if slugOptional == nil {
		slug = GenerateSlug(title)
	} else {
		slug = *slugOptional
	}

	if s.repository.GetBySlug(ctx, slug) != nil {
		return nil, errors.New("Slug error")
	}

	post := model.NewPost(userId, title, slug, content)
	err := s.repository.Save(ctx, &post)

	return &post, err
}
