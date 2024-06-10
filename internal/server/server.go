package server

import (
	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"

	_ "stl-file-analysis/docs"

	stlttp "stl-file-analysis/internal/stl/delivery/http"
	"stl-file-analysis/internal/stl/usecase"
)

func (s *Server) MapControllers(app *echo.Echo) error {

	app.GET("/swagger/*", echoswagger.WrapHandler)

	// UseCases
	stlUseCase := usecase.NewStlUseCase()

	// Controllers
	stlController := stlttp.NewStlController(stlUseCase)

	// Routes
	stlttp.MapStlRoutes(app.Group("/stl"), stlController)

	return nil
}
