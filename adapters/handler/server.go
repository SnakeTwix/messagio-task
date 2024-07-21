package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"messagio/core/enums/env"
	"messagio/utils"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

type Server struct {
	Echo *echo.Echo
}

func GetServerInstance() *Server {
	echoInstance := echo.New()
	echoInstance.IPExtractor = echo.ExtractIPFromXFFHeader()
	echoInstance.Validator = &CustomValidator{validator: validator.New()}

	server := &Server{
		Echo: echoInstance,
	}

	return server
}

func (s *Server) StartDebug() {
	s.Echo.Logger.Info(s.Echo.Start(utils.GetEnv(env.ApiAddress)))
	//s.Echo.Logger.Info(s.Echo.Start(":1234"))
}

func (s *Server) Start() {
	log.Fatal("NOT IMPLEMENTED")
}
