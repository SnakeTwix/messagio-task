package handler

import (
	"github.com/labstack/echo/v4"
	"messagio/core/entity"
	"messagio/core/ports"
	"net/http"
	"strconv"
)

type Message struct {
	serviceMessage ports.ServiceMessage
}

var handlerMessage *Message

func GetMessage(serviceMessage ports.ServiceMessage) *Message {
	if handlerMessage != nil {
		return handlerMessage
	}

	handlerMessage = &Message{
		serviceMessage: serviceMessage,
	}

	return handlerMessage
}

func (h *Message) RegisterRoutes(group *echo.Group) {
	group.POST("/message", h.createMessage)
	group.GET("/message/:id", h.getMessage)
}

func (h *Message) createMessage(ctx echo.Context) error {
	var message entity.CreateMessage

	if err := ctx.Bind(&message); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&message); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err := h.serviceMessage.CreateMessage(ctx.Request().Context(), &message)

	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.String(http.StatusCreated, message.Content)
}

func (h *Message) getMessage(ctx echo.Context) error {
	messageId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Id needs to be a number")
	}

	message, err := h.serviceMessage.GetMessage(ctx.Request().Context(), messageId)

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, &message)

}