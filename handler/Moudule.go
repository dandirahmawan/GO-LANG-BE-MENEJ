package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

type ResultModule struct {
	DataModule      []model.SectionModel `json:"dataModule"`
	DataProjectTeam interface{}          `json:"dataProjectTeam"`
	LabelsList      []model.LabelsModel  `json:"labelsList"`
}

func GetModule(ctx *gin.Context) {
	id := ctx.Param("id")

	dataModule := services.GetModuleByProjecId(id)
	dataTeam := services.GetProjectTeamByProjectId(ctx)
	dataLabels := services.GetLabelsByProjectId(ctx)

	/*inisiasi response data*/
	var data ResultModule
	data.DataModule = dataModule
	data.DataProjectTeam = dataTeam
	data.LabelsList = dataLabels
	// data := services.GetModuleByProjecId(int64(param))
	ctx.IndentedJSON(http.StatusOK, data)
}
