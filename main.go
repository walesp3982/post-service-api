package main

import (
	"context"
	"walesp3982/golang-post-api/config"
	"walesp3982/golang-post-api/model"
	"walesp3982/golang-post-api/pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBConnection(url string) *gorm.DB {
	pkg.GetLogger().Info("Access to the db")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		pkg.GetLogger().Error("Cannot access to db")
		pkg.GetLogger().Error(err.Error())
		pkg.GetLogger().Error("Aborting...")
		return nil
	}
	return db
}

func main() {
	pkg.GetLogger().Info("Gettings variables environment")
	config := config.New()

	db := getDBConnection(config.DatabaseUrl)
	if db == nil {
		return
	}

	pkg.GetLogger().Info("Migrating db")

	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.RefreshToken{}); err != nil {
		pkg.GetLogger().Warn("Error into migrate db")
		pkg.GetLogger().Warn(err.Error())
	}
	user := model.NewUser("juan", "j@gmail.com", "password")

	ctx := context.Background()

	pkg.GetLogger().Info("Creating user")
	err := gorm.G[model.User](db).Create(ctx, &user)

	if err != nil {
		pkg.GetLogger().Error("Cannot create user")
		pkg.GetLogger().Error(err.Error())
		return
	}
	pkg.GetLogger().Info("User created")
}
