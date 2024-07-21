package message

import (
	"github.com/labstack/echo/v4"
	"messagio/core/ports"
)

type Handler struct {
	serviceMessage ports.ServiceMessage
}

var handlerMessage *Handler

func GetHandler(serviceMessage ports.ServiceMessage) *Handler {
	if handlerMessage != nil {
		return handlerMessage
	}

	handlerMessage = &Handler{
		serviceMessage: serviceMessage,
	}

	return handlerMessage
}

func (h *Handler) RegisterRoutes(group *echo.Group) {
	group.POST("/message", h.createMessage)
}
