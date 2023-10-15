package router

import (
	"github.com/dandirahmawan/menej_api_go/handler"
	"github.com/dandirahmawan/menej_api_go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	route := gin.New()
	route.MaxMultipartMemory = 10 << 20 // 10 MiB

	route.Use(middleware.CORSMiddleware())
	authorizedApi := route.Use(middleware.AuthMiddleware(db))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authorizedApi.GET("/project/list", handler.GetAllProjects)
	authorizedApi.GET("/module/data/:id", handler.GetModule)
	authorizedApi.GET("/start_data", handler.GetStartData)
	authorizedApi.GET("/module/detail/:id", handler.GetModuleById)
	authorizedApi.GET("/module/assigning/:id", handler.GetModuleAssigning)
	authorizedApi.GET("/attachment/list/:id", handler.GetDocumentFile)
	authorizedApi.GET("/permition/:projectid/:userid", handler.GetPermitionProject)
	authorizedApi.GET("/invitation/:id", handler.Invitation)
	authorizedApi.GET("/file/:id", handler.DownloadFile)

	authorizedApi.POST("/project", handler.PostProject)
	authorizedApi.POST("/module", handler.PostModule)
	authorizedApi.POST("/project/list", handler.GetAllProjects)
	authorizedApi.POST("/attachment", handler.PostDocumentFile)
	authorizedApi.POST("/label", handler.PostLabel)
	authorizedApi.POST("/section", handler.PostSection)
	authorizedApi.POST("/set_permition", handler.SetPermition)
	authorizedApi.POST("/invite", handler.InviteMember)
	authorizedApi.POST("/delete_label", handler.DeleteLabel)

	authorizedApi.PUT("/module/:id", handler.PutModule)

	authorizedApi.DELETE("/section/:id", handler.DeleteSection)
	authorizedApi.DELETE("/member/:projectId/:userId", handler.DeleteMember)
	authorizedApi.DELETE("/attachment/:id", handler.DeleteDocumentFile)
	authorizedApi.DELETE("/module/:id", handler.DeleteModule)

	route.POST("/loginApp", handler.Login)
	route.POST("/register", handler.Register)
	return route
}
