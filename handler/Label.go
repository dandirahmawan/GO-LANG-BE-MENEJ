package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

func PostLabel(ctx *gin.Context) {
	data := services.SaveDataLabel(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}
