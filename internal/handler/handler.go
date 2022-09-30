package handler

import (
	"github.com/gustavolimam/coinext-challenge/internal/services"
	"github.com/labstack/echo/v4"
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

	return h.UserService.CreateUser()
}

func (h *handler) AddOrRemoveItem(c echo.Context) error {

	return h.UserService.AddOrRemoveItem()
}

func (h *handler) Trade(c echo.Context) error {
	return h.UserService.Trade()
}
