package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id          uuid.UUID      `gorm:"type:uuid" json:"id"`
	UserId      uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Title       string         `gorm:"type:varchar(255);not null;index:idx_title" json:"title"`
	Slug        string         `gorm:"uniqueIndex;type:varchar(255);not null;idx_slug" json:"slug"`
	Content     string         `json:"content"`
	PublishedAt time.Time      `gorm:"index:idx_published" json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at"`
	User        *User          `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func NewPost(userId uuid.UUID, title string, slug string, content string) Post {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return Post{
		Id:      uuid,
		UserId:  userId,
		Title:   title,
		Slug:    slug,
		Content: content,
	}

}
