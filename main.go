package main

import (
	"GO_AUTH/connectors"
	"GO_AUTH/initializers"
	"GO_AUTH/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()
	router.POST("/register", connectors.Register)
	router.POST("/login", connectors.Login)
	router.GET("/validate", middleware.AuthMiddleware, connectors.Validate)

	router.Run()
}
