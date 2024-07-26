package handler

import (
	"github.com/labstack/echo/v4"
	"messagio/core/entity"
	"messagio/core/ports"
	"net/http"
	"strconv"
)

type Statistic struct {
	serviceMessage ports.ServiceMessage
}

var handlerStatistic *Statistic

func GetStatistic(serviceMessage ports.ServiceMessage) *Statistic {
	if handlerStatistic != nil {
		return handlerStatistic
	}

	handlerStatistic = &Statistic{
		serviceMessage: serviceMessage,
	}

	return handlerStatistic
}

func (h *Statistic) RegisterRoutes(group *echo.Group) {
	group.GET("/statistic/message/:id", h.GetMessageStatistic)
}

func (h *Statistic) GetMessageStatistic(ctx echo.Context) error {
	messageId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Id needs to be a number")
	}

	message, err := h.serviceMessage.GetFullMessage(ctx.Request().Context(), messageId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Something went wrong")
	}

	messageStatistic := entity.MessageStatistic{
		ID:             message.ID,
		ReadTimes:      message.ReadTimes,
		KafkaProcessed: message.KafkaProcessed,
		CreatedAt:      message.CreatedAt,
	}

	return ctx.JSON(http.StatusOK, messageStatistic)
}
