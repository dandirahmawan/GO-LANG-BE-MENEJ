package model

import (
	"time"
)

type DocumentFile struct {
	Id              string
	FileName        string
	DescriptionFile string
	FileSize        int64
	UserId          string
	Extension       string
	UploadDate      time.Time
	Path            string
	File            string
	ModulId         string
	ProjectId       string
}

func (m DocumentFile) TableName() string {
	return "document_file"
}

func (m DocumentFile) Save() {

	DB.Save(&m)
}

func (m DocumentFile) FindById() DocumentFile {

	var data DocumentFile
	DB.Where(DocumentFile{Id: m.Id}).Find(&data)
	return data
}

func (m DocumentFile) DeleteById() {

	DB.Delete(DocumentFile{Id: m.Id})
}
