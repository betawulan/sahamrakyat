package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/service"
)

type orderItemDelivery struct {
	orderItemService service.OrderItemService
}

func (o orderItemDelivery) create(c echo.Context) error {
	var orderItem model.OrderItem

	err := c.Bind(&orderItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	orderItem, err = o.orderItemService.Create(c.Request().Context(), orderItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, orderItem)
}

func RegisterOrderItemRoute(orderItemService service.OrderItemService, e *echo.Echo) {
	handler := orderItemDelivery{
		orderItemService: orderItemService,
	}

	e.POST("/item", handler.create)
}