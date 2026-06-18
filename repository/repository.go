package repository

import "gorm.io/gorm"

type Repository struct {
	User         UserRepository
	RefreshToken RefreshTokenRepository
	Post         PostRepository
}

func New(db *gorm.DB) Repository {
	return Repository{
		User:         NewDBUser(db),
		RefreshToken: NewDBRefreshToken(db),
		Post:         NewDBPost(db),
	}
}
