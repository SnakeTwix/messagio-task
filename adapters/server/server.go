package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"messagio/core/enums/env"
	"messagio/utils"
)

type Server struct {
	Echo  *echo.Echo
	Db    *gorm.DB
	Kafka *KafkaConn
}

func GetServerInstance() *Server {
	echoInstance := initEcho()
	db := initDB()
	kafka := initKafka()

	server := &Server{
		Echo:  echoInstance,
		Db:    db,
		Kafka: kafka,
	}

	return server
}

func (s *Server) Close() {
	err := s.Kafka.MessageReader.Close()
	if err != nil {
		log.Println("Couldn't close Kafka message reader: ", err)
	}

	err = s.Kafka.MessageWriter.Close()
	if err != nil {
		log.Println("Couldn't close Kafka message writer: ", err)
	}
}

func (s *Server) Start() {
	s.Echo.Logger.Info(s.Echo.Start(utils.GetEnv(env.ApiAddress)))
}
