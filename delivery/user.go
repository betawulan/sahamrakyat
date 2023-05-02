package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/service"
)

type userDelivery struct {
	userService service.UserService
}

func (u userDelivery) create(c echo.Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err = u.userService.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (u userDelivery) readByID(c echo.Context) error {
	var user model.User

	ID := c.Param("id")
	if ID != "" {
		IDint, err := strconv.Atoi(ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		user.ID = int64(IDint)
	}

	usr, err := u.userService.ReadByID(c.Request().Context(), user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, usr)
}

func (u userDelivery) update(c echo.Context) error {
	var user model.User

	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = u.userService.Update(c.Request().Context(), int64(IDint), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (u userDelivery) read(c echo.Context) error {
	filter := model.UserFilter{}

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

	users, err := u.userService.Read(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (u userDelivery) publish(c echo.Context) error {
	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	err = u.userService.Publish(c.Request().Context(), int64(IDint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (u userDelivery) unPublish(c echo.Context) error {
	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	err = u.userService.UnPublish(c.Request().Context(), int64(IDint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func RegisterUserRoute(userService service.UserService, e *echo.Echo) {
	handler := userDelivery{
		userService: userService,
	}

	e.POST("/user", handler.create)
	e.GET("/user/:id", handler.readByID)
	e.PUT("/user/:id", handler.update)
	e.GET("/user", handler.read)
	e.PATCH("/user/:id/publish", handler.publish)
	e.PATCH("/user/:id/unpublish", handler.unPublish)
}
