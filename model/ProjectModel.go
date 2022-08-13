package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
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

func GetAllProjects() string {
	return "testing"
}

func FindAll() []ProjectModel {
	db, _ := config.ConnectDB()
	var data []ProjectModel
	db.Find(&data)

	return data
}
