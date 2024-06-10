package main

import (
	"go.uber.org/zap"
	"stl-file-analysis/internal/server"
)

// @title STL File Analysis API
// @version 1.0
// @description This is a sample server for analyzing STL files.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	s := server.NewServer()
	if err = s.Run(); err != nil {
		zap.L().Fatal("Server error", zap.Error(err))
	}
}
