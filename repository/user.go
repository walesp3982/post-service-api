package repository

import (
	"context"
	"walesp3982/golang-post-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBUser struct {
	db *gorm.DB
}

func NewDBUser(db *gorm.DB) *DBUser {
	return &DBUser{
		db: db,
	}
}

func (u *DBUser) Save(ctx context.Context, user *model.User) error {
	return gorm.G[model.User](u.db).Create(ctx, user)
}

func (u *DBUser) GetById(ctx context.Context, id uuid.UUID) *model.User {
	user, err := gorm.G[model.User](u.db).Where("id = ?", id.String()).First(ctx)
	if err != nil {
		return nil
	}
	return &user
}

func (u *DBUser) GetAll(ctx context.Context) []model.User {
	users, _ := gorm.G[model.User](u.db).Find(ctx)
	return users
}

func (u *DBUser) Update(ctx context.Context, user *model.User) error {
	_, err := gorm.G[model.User](u.db).Updates(ctx, *user)
	return err
}

func (u *DBUser) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := gorm.G[model.User](u.db).Where("id = ?", id.String()).Delete(ctx)
	return err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return NewDBUser(db)
}
