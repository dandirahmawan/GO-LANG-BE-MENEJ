package handler

import (
	"fmt"
	"net/http"

	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

type PermitionResponse struct {
	DataUser         interface{} `json:"dataUser"`
	PermitionProject interface{} `json:"dataPermition"`
}

type FormPermition struct {
	PermitionCode string `form:"permition_code"`
	ProjectId     string `form:"project_id"`
	UserId        string `form:"user_id"`
}

func GetPermitionProject(ctx *gin.Context) {
	projectid := ctx.Param("projectid")
	userid := ctx.Param("userid")

	dataPermition := services.GetDataPermition(userid, projectid)

	var um model.UserModel
	um.UserId = userid
	dataUser := um.FindById()

	var resp PermitionResponse
	resp.DataUser = dataUser
	resp.PermitionProject = dataPermition
	ctx.IndentedJSON(http.StatusOK, resp)
}

func SetPermition(ctx *gin.Context) {
	var form FormPermition
	ctx.Bind(&form)
	fmt.Println(form)
	services.SetPermition(form.PermitionCode, form.ProjectId, form.UserId)
	ctx.IndentedJSON(http.StatusOK, "success")
}
