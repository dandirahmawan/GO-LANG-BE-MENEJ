package services

import (
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
		rsp.Code = 200
		rsp.Message = "User not registered"

		rspList = append(rspList, rsp)
		return rspList
	}
	return true
}

func FindAllUser() []model.UserModel {
	db, _ := config.ConnectDB()
	var data []model.UserModel

	db.Find(&data)
	return data
}
