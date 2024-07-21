package message

import (
	"github.com/labstack/echo/v4"
	"messagio/core/entity"
	"net/http"
)

func (h *Handler) createMessage(ctx echo.Context) error {
	var message entity.CreateMessage

	// TODO: Add proper validation
	if err := ctx.Bind(&message); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if len(message.Content) == 0 {
		return ctx.String(http.StatusBadRequest, "Content is required")
	}

	err := h.serviceMessage.CreateMessage(ctx.Request().Context(), &message)

	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.String(http.StatusCreated, message.Content)
}
