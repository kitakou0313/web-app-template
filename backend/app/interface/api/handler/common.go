package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
	message string
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{
		message: "Hello World 3",
	}
}

func (h *HelloHandler) HelloHandler(c echo.Context) error {
	return c.String(
		http.StatusOK, h.message,
	)
}
