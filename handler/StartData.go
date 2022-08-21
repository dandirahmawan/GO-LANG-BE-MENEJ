package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func GetStartData(ctx *gin.Context) {
	data := services.GetStartData(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}
