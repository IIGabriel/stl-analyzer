package usecase

import (
	"io"
	"stl-file-analysis/internal/stl"
	"stl-file-analysis/pkg/stlanalyzer"
	"stl-file-analysis/pkg/stlanalyzer/ascii"
)

type stlUseCase struct {
}

func NewStlUseCase() stl.UseCase {
	return &stlUseCase{}
}

func (s *stlUseCase) InitializeModel3D(file io.ReadSeeker) (stlanalyzer.Model3D, error) {
	return ascii.NewAsciiAnalyzer(file)
}

func (s *stlUseCase) CalculateSurfaceArea(triangles []stlanalyzer.Triangle) float64 {
	var totalArea float64
	for _, triangle := range triangles {
		totalArea += stlanalyzer.CalculateTriangleArea(triangle)
	}
	return totalArea
}
