package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Gettings variables environment")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env")
	}

	dsn := os.Getenv("DATABASE_URL")

	logger.Info("Access to the db")
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error("Cannot access to db")
		logger.Error(err.Error())
		logger.Error("Aborting...")
		return
	}

}
