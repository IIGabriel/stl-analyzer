package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

// NewServer New Server constructor
func NewServer() *Server {
	server := echo.New()

	return &Server{echo: server}
}

func (s *Server) Run() error {
	if err := s.MapControllers(s.echo); err != nil {
		return err
	}

	return s.echo.Start(":8080")
}
