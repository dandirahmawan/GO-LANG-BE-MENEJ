package model

import (
	"fmt"

	"github.com/dandirahmawan/menej_api_go/config"
)

type SectionModel struct {
	Id        string       `json:"id" gorm:"primaryKey"`
	ProjectId string       `json:"projectId"`
	Section   string       `json:"section"`
	Modules   []ViewModule `json:"modules" gorm:"foreignKey:SectionId;references:Id"`
}

func (s SectionModel) TableName() string {
	return "section"
}

func FindSectionByProjectId(projectId string) []SectionModel {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

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

func FindSectionOnlyByProjectId(projectId string) []SectionModel {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	var data []SectionModel
	err := db.Model(&SectionModel{}).Where(&SectionModel{ProjectId: projectId}).Find(&data)

	// db.Where(&Mod{EmailUser: email, UserPassword: pass}).Find(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
