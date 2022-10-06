package handler

import (
	"fmt"
	"net/http"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/input"
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

func PostProject(ctx *gin.Context) {
	sessionid := ctx.Request.Header.Get("sessionid")
	dataSession := commons.GetDataSession(sessionid)
	var input input.InputProject
	ctx.BindJSON(&input)
	fmt.Println(input)
	services.SaveProject(input, dataSession.AccountId)
}
