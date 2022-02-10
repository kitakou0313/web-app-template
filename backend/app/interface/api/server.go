package api

import (
	"strconv"

	"backend/app/interface/api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	server *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	s.server = echo.New()

	s.server.Debug = true

	s.setCommonMiddleware()
	// Adding Handler
	s.Route()
	return nil
}

func (s *Server) setCommonMiddleware() {
	s.server.Use(middleware.Logger())
}

func (s *Server) Run(port int) {
	if err := s.server.Start(":" + strconv.Itoa(port)); err != nil {
		return
	}
}

func (s *Server) RunTLS(port int) {
	if err := s.server.StartAutoTLS(":" + strconv.Itoa(port)); err != nil {
		return
	}
}

// config TLS certificates and redirect from http to https
func (s *Server) SetTLSConfig() {
	autocert.HostWhitelist()
	s.server.AutoTLSManager.HostPolicy = autocert.HostWhitelist("localhost")
	s.server.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
}

func (s *Server) Route() {
	helloHandler := handler.NewHelloHandler()
	s.server.GET(
		"/", helloHandler.HelloHandler,
	)
}
