package model

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
)

type SessionModel struct {
	Id          string    `json:"id"`
	AccountId   string    `json:"accountId"`
	ExpiredDate time.Time `json:"expiredDate"`
}

func (m SessionModel) TableName() string {
	return "session"
}

func SaveSession(s SessionModel) {
	s.SaveData()
}

func (m SessionModel) SaveData() {
	db, _ := config.ConnectDB()
	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	db.Create(m) /*save to db*/
}
