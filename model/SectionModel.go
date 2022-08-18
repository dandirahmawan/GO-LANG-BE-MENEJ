package model

import (
	"fmt"

	"github.com/dandirahmawan/menej_api_go/config"
)

type SectionModel struct {
	Id        int64        `json:"id" gorm:"primaryKey"`
	ProjectId int64        `json:"projectId"`
	Section   string       `json:"section"`
	Modules   []ViewModule `json:"modules" gorm:"foreignKey:SectionId;references:Id"`
}

func (s SectionModel) TableName() string {
	return "section"
}

func FindSectionByProjectId(projectId int64) []SectionModel {
	db, _ := config.ConnectDB()
	var data []SectionModel
	err := db.Model(&SectionModel{}).
		Preload("Modules").
		Preload("Modules.Label").
		Preload("Modules.AssignTo").
		Where(&SectionModel{ProjectId: projectId}).Find(&data)

	// db.Where(&Mod{EmailUser: email, UserPassword: pass}).Find(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
