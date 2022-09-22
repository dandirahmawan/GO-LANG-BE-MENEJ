package model

import (
	"time"
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
	type M ViewDocumentFile
	var data []ViewDocumentFile
	DB.Where(M{ModulId: m.ModulId}).Find(&data)
	return data
}

func (m ViewDocumentFile) FindById() ViewDocumentFile {
	type M ViewDocumentFile
	var data ViewDocumentFile

	DB.Where(ViewDocumentFile{Id: m.Id}).Find(&data)
	return data
}

func (m ViewDocumentFile) FindByProjectId() []ViewDocumentFile {
	type M ViewDocumentFile
	var data []ViewDocumentFile
	DB.Where(ViewDocumentFile{ProjectId: m.ProjectId}).Find(&data)
	return data
}
