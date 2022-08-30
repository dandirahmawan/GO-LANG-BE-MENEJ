package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

type ResultModule struct {
	DataModule      []model.SectionModel     `json:"dataModule"`
	DataProject     []model.ViewProjectModel `json:"dataProject"`
	DataProjectTeam interface{}              `json:"dataProjectTeam"`
	LabelsList      []model.LabelsModel      `json:"labelsList"`
	DataPermition   interface{}              `json:"permitionProjects"`
	DataSection     interface{}              `json:"sectionModules"`
	DataStatus      interface{}              `json:"statusModules"`
}

func PostModule(ctx *gin.Context) {
	data := services.SaveModule(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func PutModule(ctx *gin.Context) {
	data := services.SaveModule(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func GetModule(ctx *gin.Context) {
	id := ctx.Param("id")

	dataModule := services.GetModuleByProjecId(id)
	dataTeam := services.GetProjectTeamByProjectId(ctx)
	dataLabels := services.GetLabelsByProjectId(ctx)

	dataProject := services.GetProjectById(ctx)
	dataProjectArr := []model.ViewProjectModel{dataProject}

	dataPermition := services.GetDataPermition(ctx)
	dataSection := services.GetDataSection(ctx)
	dataStatus := services.GetDataStatus(ctx)

	/*inisiasi response data*/
	var data ResultModule
	data.DataModule = dataModule
	data.DataProjectTeam = dataTeam
	data.LabelsList = dataLabels
	data.DataProject = dataProjectArr
	data.DataPermition = dataPermition
	data.DataSection = dataSection
	data.DataStatus = dataStatus

	type resp struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	var response resp
	response.Success = true
	response.Data = data
	// data := services.GetModuleByProjecId(int64(param))
	ctx.IndentedJSON(http.StatusOK, response)
}

func GetModuleById(ctx *gin.Context) {
	data := services.GetDataModuleDetail(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}

func GetModuleAssigning(ctx *gin.Context) {
	data := services.GetProjectTeamByProjectId(ctx)
	ctx.IndentedJSON(http.StatusOK, data)
}
