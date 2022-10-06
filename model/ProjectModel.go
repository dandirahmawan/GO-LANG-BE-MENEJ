package model

import (
	"time"
)

type ProjectModel struct {
	ProjectId   string
	ProjectName string
	CreatedBy   string
	CreatedDate time.Time
	Pic         string
	IsClose     string
	IsDelete    string
}

type Table interface {
	TableName() string
}

func (p ProjectModel) TableName() string {
	return "project"
}

func FindAllProject() []ProjectModel {

	var data []ProjectModel
	DB.Find(&data)

	return data
}

func (m ProjectModel) Save() {
	DB.Create(m)
}
