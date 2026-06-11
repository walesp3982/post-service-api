package model

import (
	"time"
	"walesp3982/golang-post-api/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id             uuid.UUID
	Name           string
	Email          string `gorm:"uniqueIndex"`
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Posts          []Post
	RefreshToken   []RefreshToken
}

func (u *User) VerifyPassword(password string) bool {
	return pkg.VerifyPassword(u.HashedPassword, password)
}

func NewUser(name string, email string, password string) User {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	hashed_password := pkg.HashPassword(password)
	return User{
		Id:             uuid,
		Name:           name,
		Email:          email,
		HashedPassword: string(hashed_password),
	}
}
