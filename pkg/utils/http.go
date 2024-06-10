package utils

import (
	"github.com/labstack/echo/v4"
)

type HTTPResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func HttpResponse(c echo.Context, statusCode int, data any, message ...string) error {
	httpResponse := HTTPResponse{
		Data:   data,
		Status: "success",
	}
	if len(message) > 0 {
		httpResponse.Message = message[0]
	}

	if statusCode >= 300 {
		httpResponse.Status = "error"
	}
	return c.JSON(statusCode, httpResponse)
}
