package stl

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Triangles(ctx echo.Context) error
}
