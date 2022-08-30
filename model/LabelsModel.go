package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
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
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Save(m)
}

func FindLabelsByProjectId(projectId string) []LabelsModel {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	var data []LabelsModel
	db.Where(&LabelsModel{ProjectId: projectId}).Find(&data)
	return data
}
