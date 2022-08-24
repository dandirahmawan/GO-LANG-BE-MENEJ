package services

import (
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/config"
	"github.com/dandirahmawan/menej_api_go/model"
)

func FindByEmailAndPassword(email string, pass string) interface{} {
	db, _ := config.ConnectDB()
	var data []model.UserModel
	type Mod model.UserModel

	db.Where(&Mod{EmailUser: email, UserPassword: pass}).Find(&data)

	count := len(data)
	if count <= 0 {
		type resp struct {
			Code    int16  `json:"code"`
			Message string `json:"message"`
		}

		var rsp resp
		var rspList []resp
		rsp.Code = 201
		rsp.Message = "User not registered"

		rspList = append(rspList, rsp)
		return rspList
	}

	/*response success login*/
	type s struct {
		Code      int16  `json:"code"`
		SessionId string `json:"sessionId"`
		UserId    string `json:"userId"`
	}

	accountId := data[0].UserId
	token := commons.RandStringRunes(30)

	/*sava data session*/
	SaveSessionLogin(accountId, token)

	var r s
	r.Code = 200
	r.SessionId = token
	r.UserId = accountId

	/*arra respones login*/
	var arr []s
	arr = append(arr, r)

	return arr
}

func FindAllUser() []model.UserModel {
	db, _ := config.ConnectDB()
	var data []model.UserModel

	db.Find(&data)
	return data
}

func SaveSessionLogin(accountId string, token string) {
	now := time.Now()
	exp := now.AddDate(1, 0, 0)

	var sm model.SessionModel
	sm.AccountId = accountId
	sm.ExpiredDate = exp
	sm.Id = token

	model.SaveSession(sm)
}
