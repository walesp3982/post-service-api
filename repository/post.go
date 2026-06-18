package repository

import (
	"context"
	"walesp3982/golang-post-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBTask struct {
	db *gorm.DB
}

func NewDBTask(db *gorm.DB) *DBTask {
	return &DBTask{
		db: db,
	}
}

func (d *DBTask) Save(ctx context.Context, post *model.Post) error {
	err := gorm.G[model.Post](d.db).Create(ctx, post)
	return err
}

func (d *DBTask) GetById(ctx context.Context, id uuid.UUID) *model.Post {
	post, err := gorm.G[model.Post](d.db).Preload("User", nil).Where("id = ?", id.String()).First(ctx)
	if err != nil {
		return nil
	}
	return &post
}

func (d *DBTask) GetByUser(ctx context.Context, userId uuid.UUID) []model.Post {
	posts, err := gorm.G[model.Post](d.db).Where("user_id = ?", userId).Find(ctx)
	if err != nil {
		return []model.Post{}
	}
	return posts
}

func (d *DBTask) Update(ctx context.Context, post *model.Post) error {
	_, err := gorm.G[model.Post](d.db).Updates(ctx, *post)
	return err
}

func (d *DBTask) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := gorm.G[model.Post](d.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (d *DBTask) GetAllLastest(ctx context.Context, offset uint, limit uint) []model.Post {
	posts, err := gorm.G[model.Post](d.db).Order("").Offset(int(offset)).Limit(int(limit)).Find(ctx)
	if err != nil {
		return []model.Post{}
	}
	return posts
}

func (d *DBTask) GetAllByTitle(ctx context.Context, ilike string, limit uint) []model.Post {
	param := "%" + ilike + "%"
	posts, err := gorm.G[model.Post](d.db).Where("title ILIKE ?", param).Limit(int(limit)).Find(ctx)
	if err != nil {
		return []model.Post{}
	}
	return posts
}

func (d *DBTask) GetBySlug(ctx context.Context, slug string) *model.Post {
	post, err := gorm.G[model.Post](d.db).Where("slug = ?", slug).First(ctx)
	if err != nil {
		return nil
	}
	return &post
}

func NewDBPost(db *gorm.DB) PostRepository {
	return &DBTask{
		db: db,
	}
}
