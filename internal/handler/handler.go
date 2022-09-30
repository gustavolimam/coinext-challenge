package handler

import (
	"fmt"
	"net/http"

	"github.com/gustavolimam/coinext-challenge/internal/model"
	"github.com/gustavolimam/coinext-challenge/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type HandlerI interface {
	CreateUser(c echo.Context) error
	AddOrRemoveItem(c echo.Context) error
	Trade(c echo.Context) error
}

type handler struct {
	UserService services.UserI
}

func New() HandlerI {
	return &handler{services.New()}
}

func (h *handler) CreateUser(c echo.Context) error {
	user := &model.User{}

	if err := c.Bind(user); err != nil {
		log.Error().Err(err)
		return c.JSON(http.StatusBadRequest, model.CustomError{Message: "Error trying to bind the request body"})
	}

	return h.UserService.CreateUser(user)
}

func (h *handler) AddOrRemoveItem(c echo.Context) error {
	item := &model.Inventory{}

	if err := c.Bind(item); err != nil {
		log.Error().Err(err).Msg("Error trying to bind the request body")
		return c.JSON(http.StatusBadRequest, model.CustomError{Message: "Error trying to bind the request body"})
	}
	item.UserID = c.Get("user_id").(string)

	return h.UserService.AddOrRemoveItem(item)
}

func (h *handler) Trade(c echo.Context) error {
	userID := c.Get("user_id").(string)

	if err := h.UserService.Trade(userID); err != nil {
		log.Error().Err(err).Msg("Error trying to trade items")
		return c.JSON(http.StatusBadRequest, model.CustomError{Message: fmt.Sprintln("Error trying to trade items", err)})
	}
	return nil
}
