package main

import (
	"log"
	"os"
	"walesp3982/golang-post-api/model"
	"walesp3982/golang-post-api/pkg"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env")
	}

	dsn := os.Getenv("DATABASE_URL")
	db := getDBConnection(dsn)
	if db == nil {
		return
	}

	pkg.GetLogger().Info("Migrating db")
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.RefreshToken{})

}
