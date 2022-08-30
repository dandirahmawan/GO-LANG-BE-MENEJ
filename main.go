package main

import (
	"github.com/dandirahmawan/menej_api_go/handler"
	"github.com/dandirahmawan/menej_api_go/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	// _ "github.com/santosh/gingo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Gingo Bookstore API
var DB *gorm.DB

func main() {
	route := gin.Default()

	/*init database connection*/
	// db, _ := config.ConnectDB()
	// DB = db
	route.MaxMultipartMemory = 10 << 20 // 10 MiB

	route.Use(middleware.CORSMiddleware())
	route.Use(middleware.AuthMiddleware())

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.GET("/project/list", handler.GetAllProjects)
	route.GET("/module/data/:id", handler.GetModule)
	route.GET("/start_data", handler.GetStartData)
	route.GET("/module/detail/:id", handler.GetModuleById)
	route.GET("/module/assigning/:id", handler.GetModuleAssigning)

	route.POST("/module", handler.PostModule)
	route.POST("/project/list", handler.GetAllProjects)
	route.POST("/loginApp", handler.Login)
	route.POST("/attachment", handler.PostDocumentFile)
	route.POST("/label", handler.PostLabel)

	route.PUT("/module/:id", handler.PutModule)
	route.DELETE("/attachment/:id", handler.DeleteDocumentFile)
	route.Run("localhost:8888")
}
