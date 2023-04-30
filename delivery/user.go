package delivery

import (
	"net/http"

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
