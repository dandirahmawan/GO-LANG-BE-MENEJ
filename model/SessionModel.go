package model

import (
	"time"
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

func (m SessionModel) findBySessionid() []SessionModel {
	type MS SessionModel
	var data []SessionModel
	DB.Where(MS{Id: m.Id}).Find(&data)
	return data
}

func (m SessionModel) SaveData() {
	DB.Create(m) /*save to db*/
}
