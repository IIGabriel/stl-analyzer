package usecase

import (
	"os"
	"stl-file-analysis/pkg/stlanalyzer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeModel3D(t *testing.T) {
	useCase := NewStlUseCase()

	// Test case: successful initialization
	t.Run("SuccessfulInitialization", func(t *testing.T) {
		content := []byte(`solid simplePart
facet normal 0 0 0
    outer loop
        vertex 0 0 0
        vertex 1 0 0
        vertex 1 1 1
    endloop
endfacet
facet normal 0 0 0
    outer loop
        vertex 0 0 0
        vertex 0 1 1
        vertex 1 1 1
    endloop
endfacet
endsolid simplePart`)
		tmpFile, err := os.CreateTemp("", "example.stl")
		assert.NoError(t, err)
		defer assert.NoError(t, os.Remove(tmpFile.Name()))

		_, err = tmpFile.Write(content)
		assert.NoError(t, err)
		_, err = tmpFile.Seek(0, 0)
		assert.NoError(t, err)

		model, err := useCase.InitializeModel3D(tmpFile)
		assert.NoError(t, err)
		assert.NotNil(t, model)
	})
}

func TestCalculateSurfaceArea(t *testing.T) {
	useCase := NewStlUseCase()

	// Test case: calculate surface area
	t.Run("CalculateSurfaceArea", func(t *testing.T) {
		triangles := []stlanalyzer.Triangle{
			{V1: stlanalyzer.Vector3D{}, V2: stlanalyzer.Vector3D{X: 1}, V3: stlanalyzer.Vector3D{X: 1, Y: 1, Z: 1}},
			{V1: stlanalyzer.Vector3D{}, V2: stlanalyzer.Vector3D{Y: 1, Z: 1}, V3: stlanalyzer.Vector3D{X: 1, Y: 1, Z: 1}},
			{V1: stlanalyzer.Vector3D{X: 1.342, Y: 20.2342, Z: 30.234}, V2: stlanalyzer.Vector3D{Y: 1.2345, Z: 1.2345}, V3: stlanalyzer.Vector3D{X: 1.4234, Y: 1, Z: 1}},
		}
		expectedArea := stlanalyzer.CalculateTriangleArea(triangles[0]) + stlanalyzer.CalculateTriangleArea(triangles[1]) + stlanalyzer.CalculateTriangleArea(triangles[2])
		area := useCase.CalculateSurfaceArea(triangles)
		assert.Equal(t, expectedArea, area)
	})
}
