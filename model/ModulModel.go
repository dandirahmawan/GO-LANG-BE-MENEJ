package model

import (
	"time"
)

type ModulInput struct {
	Assigned    string `json:"assigned"`
	Checklist   string `json:"checklist"`
	Date        string `json:"date"`
	Desc        string `json:"desc"`
	LabelModule string `json:"labelModule"`
	ModulId     string `json:"modulId" gorm:"primaryKey"`
	ModuleName  string `json:"moduleName"`
	Section     string `json:"section"`
	Status      string `json:"status"`
	ProjectId   string `json:"projectId"`
	CreatedBy   string `json:"createdBy"`
}

type ModulModel struct {
	ModulId      string    `json:"modulId" gorm:"primaryKey"`
	ModulName    string    `json:"modulName"`
	CreatedDate  time.Time `json:"createdDate"`
	EndDate      time.Time `json:"endDate"`
	ModulStatus  string    `json:"modulStatus"`
	Description  string    `json:"description"`
	IsTrash      string    `json:"isTrash"`
	Percentage   float64   `json:"percentage"`
	OldSectionId int64     `json:"oldSectionId"`
	ProjectId    string    `json:"projectId"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	SectionId    string    `json:"sectionId"`
	UpdatedDate  time.Time `json:"updatedDate"`
}

func (m ModulModel) TableName() string {
	return "modul"
}

func (m ModulModel) DeleteById() int16 {
	var mod ModulModel
	DB.Where("modul_id = ?", m.ModulId).Delete(&mod)
	return 1
}

func (m ModulModel) FindById() ModulModel {
	type Mod ModulModel
	var data ModulModel
	DB.Where(Mod{ModulId: m.ModulId}).Find(&data)

	return data
}

func (m ModulModel) Update() ModulModel {
	DB.Updates(&m)

	return m
}

func (m ModulModel) SaveModule() interface{} {
	DB.Save(&m)

	return m
}
