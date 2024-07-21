package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"messagio/core/enums/env"
	"messagio/core/ports"
	"messagio/utils"
)

func GetServerInstance(services *Services) *Server {
	echoInstance := echo.New()
	echoInstance.IPExtractor = echo.ExtractIPFromXFFHeader()

	server := &Server{
		Echo:     echoInstance,
		Services: services,
	}

	return server
}

type Server struct {
	Echo     *echo.Echo
	Services *Services
}

type Services struct {
	ServiceMessage ports.ServiceMessage
}

func (s *Server) StartDebug() {
	s.Echo.Logger.Info(s.Echo.Start(utils.GetEnv(env.ApiAddress)))
	//s.Echo.Logger.Info(s.Echo.Start(":1234"))
}

func (s *Server) Start() {
	log.Fatal("NOT IMPLEMENTED")
}
