package model

import (
	"time"
)

type LabelsModel struct {
	ProjectId string    `json:"projectId" form:"projectId" gorm:"primaryKey"`
	Label     string    `json:"label" form:"label" gorm:"primaryKey"`
	Color     string    `json:"color" form:"color"`
	Date      time.Time `json:"date"`
}

func (l LabelsModel) TableName() string {
	return "labels"
}

func (m LabelsModel) Save() {
	DB.Save(m)
}

func (m LabelsModel) Delete() {
	DB.Delete(&LabelsModel{ProjectId: m.ProjectId, Label: m.Label})
}

func FindLabelsByProjectId(projectId string) []LabelsModel {
	var data []LabelsModel
	DB.Where(&LabelsModel{ProjectId: projectId}).Find(&data)
	return data
}

func FindLabelsByModuleId(projectId string) []LabelsModel {
	var data []LabelsModel
	DB.Where(&LabelsModel{ProjectId: projectId}).Find(&data)
	return data
}

func FindLabelByIdAndLabel(projectId string, label string) []LabelsModel {
	var data []LabelsModel
	DB.Where(&LabelsModel{ProjectId: projectId, Label: label}).Find(&data)
	return data
}
