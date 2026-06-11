package main

import (
	"walesp3982/golang-post-api/config"
	"walesp3982/golang-post-api/database"
	"walesp3982/golang-post-api/pkg"
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

}
