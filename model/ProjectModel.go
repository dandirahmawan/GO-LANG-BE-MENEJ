package model

import (
	"time"
)

type ProjectModel struct {
	ProjectId   int64
	ProjectName string
	CreatedBy   int64
	CreatedDate time.Time
	Pic         int64
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
