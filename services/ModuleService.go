package services

import "github.com/dandirahmawan/menej_api_go/model"

func GetModuleByProjecId(projectid string) []model.SectionModel {
	var data []model.SectionModel
	data = model.FindSectionByProjectId(projectid)
	return data
}
