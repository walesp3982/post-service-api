package repository

import (
	"context"
	"walesp3982/golang-post-api/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, id uuid.UUID) *model.User
	GetAll(ctx context.Context) []model.User
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *model.RefreshToken) error
	GetByToken(ctx context.Context, token string) *model.RefreshToken
	Update(ctx context.Context, token *model.RefreshToken) error
	Delete(ctx context.Context, id uuid.UUID) *model.RefreshToken
}

type PostRepository interface {
	Save(ctx context.Context, post *model.Post) error
	GetById(ctx context.Context, id uuid.UUID) *model.Post
	GetByUser(ctx context.Context, userId uuid.UUID) []model.Post
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id int) error
	GetAllLastest(ctx context.Context, offset uint, limit uint) []model.Post
	GetAllByTitle(ctx context.Context, ilike string, limit uint) []model.Post
	GetBySlug(ctx context.Context, slug string) *model.Post
}
