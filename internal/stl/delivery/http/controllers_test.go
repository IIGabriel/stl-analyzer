package http

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	mock "stl-file-analysis/internal/stl/mocks"
	"stl-file-analysis/pkg/stlanalyzer/ascii"
	"stl-file-analysis/pkg/utils"
	"testing"
)

func TestTriangles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockUseCase(ctrl)

	stlController := NewStlController(mockUseCase)
	fileContent := []byte(`solid simplePart
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
	tempFile, err := os.CreateTemp("", "example.stl")
	assert.NoError(t, err)
	defer assert.NoError(t, os.Remove(tempFile.Name()))
	_, err = tempFile.Write(fileContent)
	assert.NoError(t, err)
	_, err = tempFile.Seek(0, 0)
	assert.NoError(t, err)

	model, err := ascii.NewAsciiAnalyzer(tempFile)
	require.NoError(t, err)

	mockUseCase.EXPECT().InitializeModel3D(gomock.Any()).Return(model, nil)

	triangles, err := model.Triangles()
	require.NoError(t, err)

	mockedSurfaceArea := 1.4142
	mockUseCase.EXPECT().CalculateSurfaceArea(triangles).Return(mockedSurfaceArea)

	// Create a multipart form file
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", tempFile.Name())
	assert.NoError(t, err)
	_, err = io.Copy(part, tempFile)
	assert.NoError(t, err)
	assert.NoError(t, writer.Close())

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/stl/triangles", &body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, rec)

	// Call the handler
	err = stlController.Triangles(ctx)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response utils.HTTPResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	data := response.Data.(map[string]any)
	assert.Equal(t, float64(len(triangles)), data[`number_of_triangles`])
	assert.Equal(t, mockedSurfaceArea, data[`surface_area`])
}
