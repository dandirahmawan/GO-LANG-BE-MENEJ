package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
)

type ViewDocumentFile struct {
	Id              string    `json:"id"`
	ModulId         string    `json:"modulId"`
	ProjectId       string    `json:"projectId"`
	UserId          string    `json:"userId"`
	FileName        string    `json:"fileName"`
	FileSize        string    `json:"fileSize"`
	Extension       string    `json:"extension"`
	Path            string    `json:"path"`
	UploadDate      time.Time `json:"uploadDate"`
	DescriptionFile string    `json:"descriptionFile"`
	ModulName       string    `json:"modulName"`
	UserIdModule    string    `json:"userIdModule"`
	ProjectName     string    `json:"projectName"`
	Pic             string    `json:"pic"`
	UserName        string    `json:"userName"`
	UrlPath         string    `json:"urlPath"`
}

func (m ViewDocumentFile) TableName() string {
	return "view_document_file"
}

func (m ViewDocumentFile) FindByModulId() []ViewDocumentFile {
	db, _ := config.ConnectDB()
	type M ViewDocumentFile
	var data []ViewDocumentFile

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Where(M{ModulId: m.ModulId}).Find(&data)
	return data
}

func (m ViewDocumentFile) FindById() ViewDocumentFile {
	db, _ := config.ConnectDB()
	type M ViewDocumentFile
	var data ViewDocumentFile

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Where(ViewDocumentFile{Id: m.Id}).Find(&data)
	return data
}
