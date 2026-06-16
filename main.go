package main

import (
	"log"
	"walesp3982/golang-post-api/config"
	"walesp3982/golang-post-api/database"
	"walesp3982/golang-post-api/handler"
	"walesp3982/golang-post-api/pkg"
	"walesp3982/golang-post-api/repository"
	"walesp3982/golang-post-api/services"
)

func main() {
	pkg.GetLogger().Info("Gettings variables environment")
	config := config.New()

	// Get DB Connection
	db, err := database.GetDBConnection(config.DatabaseUrl)
	if err != nil {
		pkg.GetLogger().Error("Cannot access db")
		pkg.GetLogger().Error(err.Error())
		return
	}

	// Run migrations
	pkg.GetLogger().Info("Running all migrations")
	database.RunMigration(db)

	repository := repository.New(db)
	services := services.New(repository, config)
	server := handler.New(services)
	log.Fatal(server.Listen(":3000"))
}
