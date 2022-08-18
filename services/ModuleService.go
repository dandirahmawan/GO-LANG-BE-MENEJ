package services

import "github.com/dandirahmawan/menej_api_go/model"

func GetModuleByProjecId(projectid int64) []model.SectionModel {
	var data []model.SectionModel
	data = model.FindSectionByProjectId(projectid)
	return data
}
