package model

import (
	"fmt"
	"time"
)

type SectionModel struct {
	Id          string       `json:"id" gorm:"primaryKey"`
	ProjectId   string       `json:"projectId"`
	Section     string       `json:"section"`
	CreatedDate time.Time    `json:"createdDate"`
	Modules     []ViewModule `json:"modules" gorm:"foreignKey:SectionId;references:Id"`
}

func (s SectionModel) TableName() string {
	return "section"
}

func (m SectionModel) Save() {
	//

	// defer func() {
	// 	config.CloseConnection(db)
	// }()

	DB.Save(m)
}

func FindSectionByProjectId(projectId string) []SectionModel {
	var data []SectionModel
	err := DB.Model(&SectionModel{}).
		Preload("Modules").
		Preload("Modules.Label").
		Preload("Modules.AssignTo").
		Where(&SectionModel{ProjectId: projectId}).
		Order("created_date ASC").
		Find(&data)

	if err != nil {
		fmt.Println(err)
	}
	return data
}

func DeleteSectionById(id string) int16 {
	var Section SectionModel
	err := DB.Where("id = ?", id).Delete(&Section).Error

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return 1
}

func FindSectionOnlyByProjectId(projectId string) []SectionModel {
	var data []SectionModel
	err := DB.Where(&SectionModel{ProjectId: projectId}).Order("created_date ASC").Find(&data)

	// fmt.Println(data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
