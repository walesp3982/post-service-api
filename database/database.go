package database

import (
	"walesp3982/golang-post-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.RefreshToken{})
	if err != nil {
		return err
	}
	return nil
}
