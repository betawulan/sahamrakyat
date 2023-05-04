package delivery

import (
	"net/http"
	"strconv"
	"strings"

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

func (o orderItemDelivery) readByID(c echo.Context) error {
	orderItem := model.OrderItem{}

	ID := c.Param("id")
	if ID != "" {
		IDint, err := strconv.Atoi(ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		orderItem.ID = int64(IDint)
	}

	item, err := o.orderItemService.ReadByID(c.Request().Context(), orderItem.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, item)
}

func (o orderItemDelivery) update(c echo.Context) error {
	var orderItem model.OrderItem

	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = c.Bind(&orderItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = o.orderItemService.Update(c.Request().Context(), int64(IDint), orderItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (o orderItemDelivery) read(c echo.Context) error {
	filter := model.OrderItemFilter{}

	limit := c.QueryParam("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		filter.Limit = uint64(limitInt)
	}

	page := c.QueryParam("page")
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		filter.Page = pageInt
	}

	status := c.QueryParam("status")
	if status != "" {
		statuses := strings.Split(status, ",")
		for _, value := range statuses {
			statusInt, err := strconv.Atoi(value)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			filter.Status = append(filter.Status, statusInt)
		}
	}

	ordersItem, err := o.orderItemService.Read(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, ordersItem)
}

func (o orderItemDelivery) publish(c echo.Context) error {
	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	err = o.orderItemService.Publish(c.Request().Context(), int64(IDint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (o orderItemDelivery) unPublish(c echo.Context) error {
	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	err = o.orderItemService.UnPublish(c.Request().Context(), int64(IDint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func RegisterOrderItemRoute(orderItemService service.OrderItemService, e *echo.Echo) {
	handler := orderItemDelivery{
		orderItemService: orderItemService,
	}

	e.POST("/item", handler.create)
	e.GET("/item/:id", handler.readByID)
	e.PUT("/item/:id", handler.update)
	e.GET("/item", handler.read)
	e.PATCH("/item/:id/publish", handler.publish)
	e.PATCH("/item/:id/unpublish", handler.unPublish)
}
