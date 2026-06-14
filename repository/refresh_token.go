package repository

import (
	"context"
	"walesp3982/golang-post-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBRefreshToken struct {
	db *gorm.DB
}

func NewDBRefreshToken(db *gorm.DB) *DBRefreshToken {
	return &DBRefreshToken{
		db: db,
	}
}

func (r *DBRefreshToken) Save(ctx context.Context, token *model.RefreshToken) error {
	return gorm.G[model.RefreshToken](r.db).Create(ctx, token)
}

func (r *DBRefreshToken) GetByToken(ctx context.Context, token string) *model.RefreshToken {
	tokenDb, err := gorm.G[model.RefreshToken](r.db).Where("token = ?", token).First(ctx)
	if err != nil {
		return nil
	}
	return &tokenDb
}

func (r *DBRefreshToken) Update(ctx context.Context, token *model.RefreshToken) error {
	_, err := gorm.G[model.RefreshToken](r.db).Updates(ctx, *token)
	return err
}

func (r *DBRefreshToken) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := gorm.G[model.RefreshToken](r.db).Where("id = ?", id.String()).Delete(ctx)
	return err
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return NewDBRefreshToken(db)
}
