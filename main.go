package main

import (
	"github.com/dandirahmawan/menej_api_go/handler"
	"github.com/dandirahmawan/menej_api_go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Use(middleware.CORSMiddleware())
	route.Use(middleware.AuthMiddleware())

	route.GET("/project/list", handler.GetAllProjects)
	route.POST("/project/list", handler.GetAllProjects)
	route.POST("/login", handler.Login)
	route.POST("/module/data/:id", handler.GetModule)
	route.Run("localhost:8888")
}
