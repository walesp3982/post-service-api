package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Title       string
	Slug        string `gorm:"uniqueIndex"`
	Content     string
	PublishedAt time.Time
	CreatedAt   time.Time
	DeleteAt    gorm.DeletedAt
	Status      string
	User        User `gorm:"constraint:OnDelete:CASCADE"`
}
