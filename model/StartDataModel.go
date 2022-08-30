package model

import (
	"fmt"

	"github.com/dandirahmawan/menej_api_go/config"
)

type StartData struct {
	Email            string `json:"email"`
	Name             string `json:"name"`
	PicProfile       string `json:"picProfile"`
	PicProfileDetail string `json:"picProfileDetail"`
	UnreadNotif      int32  `json:"unreadNotif"`
}

func FindDataStart(sessionid string) StartData {
	SQL := "select \n" +
		"	usr.email_user email, \n" +
		"	usr.user_name name, \n" +
		"	usr.pic_profile picProfile, \n" +
		"	usr.pic_profile_detail picProfileDetail, \n" +
		"	0 unreadNotif \n" +
		"from user usr \n" +
		"join session ssn on ssn.account_id = usr.user_id \n" +
		"where ssn.id = '" + sessionid + "'"

	fmt.Println(SQL)
	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	var data StartData
	db.Raw(SQL).Find(&data)
	return data
}
