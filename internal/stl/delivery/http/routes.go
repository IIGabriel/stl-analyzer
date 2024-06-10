package http

import (
	"github.com/labstack/echo/v4"
	"stl-file-analysis/internal/stl"
)

func MapStlRoutes(group *echo.Group, h stl.Controller) {
	group.POST("/triangles", h.Triangles)
}
