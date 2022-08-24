package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
)

type LabelsModel struct {
	ProjectId string    `json:"projectId" gorm:"primaryKey"`
	Label     string    `json:"label" gorm:"primaryKey"`
	Color     string    `json:"color"`
	Date      time.Time `json:"date"`
}

func (l LabelsModel) TableName() string {
	return "labels"
}

func FindLabelsByProjectId(projectId string) []LabelsModel {
	db, _ := config.ConnectDB()

	var data []LabelsModel
	db.Where(&LabelsModel{ProjectId: projectId}).Find(&data)
	return data
}
