package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetAllProjects(ctx *gin.Context) {
	var resp interface{}
	resp = services.GetAllViewProjects(ctx)
	ctx.IndentedJSON(http.StatusOK, resp)
}
