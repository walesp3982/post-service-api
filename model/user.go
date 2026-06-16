package model

import (
	"time"
	"walesp3982/golang-post-api/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id             uuid.UUID      `gorm:"type:uuid" json:"id"`
	Name           string         `gorm:"type:varchar(50);not null" json:"name"`
	Email          string         `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	HashedPassword string         `gorm:"uniqueIndex;type:varchar(255);not null" json:"-"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Posts          []Post         `json:"-"`
	RefreshToken   []RefreshToken `json:"-"`
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
