package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stl-file-analysis/internal/models"
	"stl-file-analysis/internal/stl"
	"stl-file-analysis/pkg/utils"
)

type stlController struct {
	uc stl.UseCase
}

func NewStlController(sessionUC stl.UseCase) stl.Controller {
	return &stlController{sessionUC}
}

// Triangles handles the POST /triangles endpoint
// @Summary Get number of triangles and surface area of an STL file
// @Description Receives an STL file and returns the number of triangles and the surface area
// @Tags stl
// @Accept mpfd
// @Produce json
// @Param file formData file true "STL file"
// @Success 200 {object} utils.HTTPResponse{data=models.TrianglesHTTPResponse}
// @Failure 400 {object} utils.HTTPResponse{data=interface{}}
// @Failure 500 {object} utils.HTTPResponse
// @Router /stl/triangles [post]
func (s *stlController) Triangles(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return utils.HttpResponse(ctx, http.StatusBadRequest, err, "failed to get the file")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return utils.HttpResponse(ctx, http.StatusBadRequest, err, "failed to open the file")
	}
	defer file.Close()
	model3d, err := s.uc.InitializeModel3D(file)
	if err != nil {
		return utils.HttpResponse(ctx, http.StatusInternalServerError, err, "failed to initialize the 3D model")
	}

	triangles, err := model3d.Triangles()
	if err != nil {
		return utils.HttpResponse(ctx, http.StatusInternalServerError, err, "failed to get triangles from the model")
	}

	areaOfModel3d := s.uc.CalculateSurfaceArea(triangles)
	return utils.HttpResponse(ctx, http.StatusOK, models.TrianglesHTTPResponse{
		NumberOfTriangles: uint64(len(triangles)),
		SurfaceArea:       areaOfModel3d,
	})
}
