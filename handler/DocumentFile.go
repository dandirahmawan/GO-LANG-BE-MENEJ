package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func PostDocumentFile(ctx *gin.Context) {
	data := services.Upload(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func GetDocumentFile(ctx *gin.Context) {
	data := services.GetDocumentFileByProjectId(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func DeleteDocumentFile(ctx *gin.Context) {
	data := services.DeleteDocumentFile(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func DownloadFile(ctx *gin.Context) {
	services.DownloadFile(ctx)
}
