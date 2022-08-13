package main

import (
	"github.com/dandirahmawan/menej_api_go/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/api/projects", handler.GetAllProjects)
	route.Run()
}
