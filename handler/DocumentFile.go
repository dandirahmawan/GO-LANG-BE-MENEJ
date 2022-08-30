package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func PostDocumentFile(ctx *gin.Context) {
	data := services.Upload(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
	// ctx.String(http.StatusOK, ("'%s' uploaded!"))
	// ctx.IndentedJSON(http.StatusOK, "dandi rahmawan")
}

func DeleteDocumentFile(ctx *gin.Context) {
	data := services.DeleteDocumentFile(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}
