package services

import (
	"strings"

	"github.com/dandirahmawan/menej_api_go/model"
)

func GetDataPermition(userid string, projectid string) interface{} {
	var data []model.PermitionProject
	data = model.FindDataPermitionByProjectIdAndUserId(userid, projectid)
	return data
}

func SetPermition(permitionCode string, projectId string, userId string) {
	/*delete old permition first*/
	var modelPermition model.PermitionModel
	modelPermition.ProjectId = projectId
	modelPermition.UserId = userId
	modelPermition.DeleteByProjectIdUserId()

	var arrp []string
	arrp = strings.Split(permitionCode, ",")
	for i := 0; i < len(arrp); i++ {
		modelPermition.PermitionCode = arrp[i]
		modelPermition.Save()
	}
}
