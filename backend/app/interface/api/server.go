package api

import (
	"strconv"

	"backend/app/interface/api/handler"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	s.server = echo.New()

	// Adding Handler
	s.Route()
	return nil
}

func (s *Server) Run(port int) {
	if err := s.server.Start(":" + strconv.Itoa(port)); err != nil {
		return
	}
}

func (s *Server) Route() {
	helloHandler := handler.NewHelloHandler()
	s.server.GET(
		"/", helloHandler.HelloHandler,
	)
}
