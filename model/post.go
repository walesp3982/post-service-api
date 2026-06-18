package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id          uuid.UUID `gorm:"type:uuid"`
	UserId      uuid.UUID `gorm:"type:uuid;not null"`
	Title       string    `gorm:"type:varchar(255);not null;index:idx_title"`
	Slug        string    `gorm:"uniqueIndex;type:varchar(255);not null;idx_slug"`
	Content     string
	PublishedAt time.Time `gorm:"index:idx_published"`
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeleteAt    gorm.DeletedAt
	User        *User `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
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
