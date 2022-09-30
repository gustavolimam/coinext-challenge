package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MustReceiveID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, "idNotReceived")
		}

		c.Set("user_id", id)

		return next(c)
	}
}
