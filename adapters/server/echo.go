package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

func initEcho() *echo.Echo {
	echoInstance := echo.New()
	echoInstance.IPExtractor = echo.ExtractIPFromXFFHeader()
	echoInstance.Validator = &CustomValidator{validator: validator.New()}

	return echoInstance
}
