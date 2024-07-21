package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"messagio/adapters/handler"
	messageHandler "messagio/adapters/handler/message"
	"messagio/adapters/repository"
	messageRepo "messagio/adapters/repository/message"
	"messagio/adapters/repository/migrations"
	messageService "messagio/core/service/message"
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
	repoMessage := messageRepo.GetRepo(db)

	// Services
	serviceMessage := messageService.GetService(repoMessage)

	// Handlers
	handlerMessage := messageHandler.GetHandler(serviceMessage)

	// Server setup
	services := handler.Services{
		ServiceMessage: serviceMessage,
	}
	server := handler.GetServerInstance(&services)
	group := server.Echo.Group("/api/v1")

	// Routes
	handlerMessage.RegisterRoutes(group)

	server.StartDebug()
}
