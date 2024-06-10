package stl

import (
	"io"
	"stl-file-analysis/pkg/stlanalyzer"
)

type UseCase interface {
	InitializeModel3D(file io.ReadSeeker) (stlanalyzer.Model3D, error)
	CalculateSurfaceArea(model []stlanalyzer.Triangle) float64
}
