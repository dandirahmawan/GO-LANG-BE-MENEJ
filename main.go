package main

import (
	"github.com/dandirahmawan/menej_api_go/handler"
	"github.com/dandirahmawan/menej_api_go/middleware"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"

	// _ "github.com/santosh/gingo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	route := gin.Default()

	/*init model connections*/
	model.ConnectDB()

	// to := []string{"dandirahmawan95@gmail.com"}
	// cc := []string{}
	// config.SendMail(to, cc, "Subject Email", config.MessageEmailInvitation("Dandi Rahmawan", "124343"))

	route.MaxMultipartMemory = 10 << 20 // 10 MiB

	route.Use(middleware.CORSMiddleware())
	// route.Use(middleware.AuthMiddleware(db))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.GET("/project/list", handler.GetAllProjects)
	route.GET("/module/data/:id", handler.GetModule)
	route.GET("/start_data", handler.GetStartData)
	route.GET("/module/detail/:id", handler.GetModuleById)
	route.GET("/module/assigning/:id", handler.GetModuleAssigning)
	route.GET("/attachment/list/:id", handler.GetDocumentFile)
	route.GET("/permition/:projectid/:userid", handler.GetPermitionProject)
	route.GET("/invitation/:id", handler.Invitation)

	route.POST("/project", handler.PostProject)
	route.POST("/module", handler.PostModule)
	route.POST("/project/list", handler.GetAllProjects)
	route.POST("/loginApp", handler.Login)
	route.POST("/attachment", handler.PostDocumentFile)
	route.POST("/label", handler.PostLabel)
	route.POST("/section", handler.PostSection)
	route.POST("/set_permition", handler.SetPermition)
	route.POST("/register", handler.Register)
	route.POST("/invite", handler.InviteMember)
	route.POST("/delete_label", handler.DeleteLabel)

	route.GET("/file/:id", handler.DownloadFile)
	route.PUT("/module/:id", handler.PutModule)
	route.DELETE("/section/:id", handler.DeleteSection)
	route.DELETE("/member/:projectId/:userId", handler.DeleteMember)
	route.DELETE("/attachment/:id", handler.DeleteDocumentFile)
	route.DELETE("/module/:id", handler.DeleteModule)
	route.Run("localhost:8888")
}
