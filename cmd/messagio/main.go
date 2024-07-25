package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"messagio/adapters/handler"
	"messagio/adapters/repository"
	"messagio/adapters/server"
	"messagio/adapters/server/migrations"
	"messagio/core/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Server setup
	server := server.GetServerInstance()
	if err := migrations.RunMigrations(server.Db); err != nil {
		log.Fatal(err)
	}
	group := server.Echo.Group("/api/v1")

	// Repos
	repoMessage := repository.GetMessage(server.Db, server.Kafka.MessageReader, server.Kafka.MessageWriter)

	// Services
	serviceMessage := service.GetMessage(repoMessage)

	// Handlers
	handlerMessage := handler.GetMessage(serviceMessage)

	// Routes
	handlerMessage.RegisterRoutes(group)

	server.StartDebug()
}
