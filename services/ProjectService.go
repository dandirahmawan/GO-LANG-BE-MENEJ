package services

import (
	"fmt"
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/input"
	"github.com/dandirahmawan/menej_api_go/model"
)

func SaveProject(input input.InputProject, accountId string) model.ProjectModel {
	var pm model.ProjectModel
	uuid := commons.GeneratdUUID(20)

	/*validate existing*/
	data := pm.FindByProjectName(input.ProjectName)
	fmt.Println(input.ProjectName)
	fmt.Println(data)
	if data.ProjectId != "" {
		pm.ProjectId = "exists"
		return pm
	} else {
		pm.ProjectId = uuid
		fmt.Println(uuid)
		pm.ProjectName = input.ProjectName
		pm.CreatedDate = time.Now()
		pm.IsClose = "N"
		pm.IsDelete = "N"
		pm.CreatedBy = accountId
		pm.Pic = accountId
		pm.Save()
		return pm
	}
}
