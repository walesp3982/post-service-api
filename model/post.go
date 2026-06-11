package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id          uuid.UUID `gorm:"type:uuid"`
	UserId      uuid.UUID `gorm:"type:uuid;not null"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Slug        string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Content     string
	PublishedAt time.Time
	CreatedAt   time.Time
	DeleteAt    gorm.DeletedAt
	User        User `gorm:"constraint:OnDelete:CASCADE"`
}
