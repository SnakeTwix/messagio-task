package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"messagio/adapters/handler"
	"messagio/adapters/repository"
	"messagio/adapters/repository/migrations"
	"messagio/core/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := repository.InitDB()
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	// Repos
	repoMessage := repository.GetMessage(db)

	// Services
	serviceMessage := service.GetMessage(repoMessage)

	// Handlers
	handlerMessage := handler.GetMessage(serviceMessage)

	// Server setup

	server := handler.GetServerInstance()
	group := server.Echo.Group("/api/v1")

	// Routes
	handlerMessage.RegisterRoutes(group)

	server.StartDebug()
}
