package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
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
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Save(&m)
}

func (m DocumentFile) FindById() DocumentFile {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	var data DocumentFile
	db.Where(DocumentFile{Id: m.Id}).Find(&data)
	return data
}

func (m DocumentFile) DeleteById() {
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Delete(DocumentFile{Id: m.Id})
}
