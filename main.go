package main

import (
	"github.com/dandirahmawan/menej_api_go/handler"
	"github.com/dandirahmawan/menej_api_go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Gingo Bookstore API
func main() {
	route := gin.Default()

	route.Use(middleware.CORSMiddleware())
	route.Use(middleware.AuthMiddleware())

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.GET("/project/list", handler.GetAllProjects)
	route.POST("/project/list", handler.GetAllProjects)
	route.POST("/login", handler.Login)
	route.POST("/module/data/:id", handler.GetModule)
	route.GET("/start_data", handler.GetStartData)
	route.Run("localhost:8888")
}
