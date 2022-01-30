package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

func hello(c echo.Context) error {
	return c.String(
		http.StatusOK, "Hello, World! 2",
	)
}

func (s *Server) Init() error {
	s.server = echo.New()

	s.server.GET("/", hello)
	return nil
}

func (s *Server) Run(port int) {
	if err := s.server.Start(":" + strconv.Itoa(port)); err != nil {
		return
	}
}
