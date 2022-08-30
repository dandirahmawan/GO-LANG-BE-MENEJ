package model

import "github.com/dandirahmawan/menej_api_go/config"

type PermitionProject struct {
	PermitionConde       int16
	PermitionName        string
	PermitionDescription string
	IsChecked            string
}

func FindDataPermitionByProjectIdAndUserId(userId string, projectId string) []PermitionProject {
	var sql string = "SELECT * FROM( \n" +
		"					SELECT \n" +
		"						pc.permition_code, \n" +
		"						pc.permition_name, \n" +
		"						pc.permition_description, \n" +
		"						'Y' is_checked \n" +
		"					FROM \n" +
		"						permition_code pc  \n" +
		"					LEFT JOIN permition p ON pc.permition_code = p.permition_code \n" +
		"					WHERE p.user_id = '" + userId + "' AND p.project_id = '" + projectId + "' \n" +
		"					UNION \n" +
		"					SELECT \n" +
		"						pc.permition_code, \n" +
		"						pc.permition_name, \n" +
		"						pc.permition_description, \n" +
		"						'N' is_checked \n" +
		"					FROM \n" +
		"						permition_code pc \n" +
		"					WHERE pc.permition_code NOT IN (SELECT \n" +
		"							pc.permition_code \n" +
		"						FROM \n" +
		"							permition_code pc  \n" +
		"						LEFT JOIN permition p ON pc.permition_code = p.permition_code \n" +
		"						WHERE p.user_id = '" + userId + "' AND p.project_id = '" + projectId + "') \n" +
		"				) R \n" +
		"				ORDER BY R.permition_code"

	db, _ := config.ConnectDB()

	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	var data []PermitionProject
	db.Raw(sql).Find(&data)
	return data
}
