package model

type PermitionProject struct {
	PermitionCode        int16  `json:"permitionCode"`
	PermitionName        string `json:"permitionName"`
	PermitionDescription string `json:"permitionDescription"`
	IsChecked            string `json:"isChecked"`
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

	var data []PermitionProject
	DB.Raw(sql).Find(&data)
	return data
}
