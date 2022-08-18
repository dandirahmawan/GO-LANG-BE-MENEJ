package model

var RawQueryViewProject string = "SELECT DISTINCT \n" +
	"	p.project_id, \n" +
	"	p.pic, \n" +
	"	p.project_name, \n" +
	"	p.created_by, \n" +
	"	DATE_FORMAT(p.created_date, '%d %M %Y') created_date, \n" +
	"	p.is_delete, \n" +
	"	p.is_close, \n" +
	"	u.user_name, \n" +
	"	(SELECT COUNT(project_id) FROM project_team WHERE project_id = p.project_id) AS count_team, \n" +
	"	u.user_name AS pic_name, \n" +
	"	u.email_user AS pic_email, \n" +
	"	(SELECT COUNT(project_id) FROM bugs WHERE project_id = p.project_id AND (is_delete = 'N' OR is_delete IS NULL)) AS count_bugs, \n" +
	"	(SELECT COUNT(modul_id) FROM modul WHERE project_id = p.project_id and (is_trash = 'N' or is_trash IS NULL)) AS count_module \n" +
	"FROM project p \n" +
	"LEFT JOIN project_team pt ON pt.project_id = p.project_id \n" +
	"LEFT JOIN User u ON p.pic = u.user_id \n" +
	"WHERE (p.is_delete <> 'Y' OR p.is_delete IS NULL)"

type ViewProjectModel struct {
	ProjectId   int64  `json:"projectId" gorm:"column:project_id"`
	Pic         int64  `json:"pic" gorm:"column:pic"`
	ProjectName string `json:"projectName" gorm:"column:project_name"`
	CreatedBy   int64  `json:"createdBy" gorm:"column:created_by"`
	CreatedDate string `json:"createdDate" gorm:"column:created_date"`
	IsDelete    string `json:"isDelete" gorm:"column:is_delete"`
	IsClose     string `json:"isClose" gorm:"column:is_close"`
	UserName    string `json:"userName" gorm:"column:user_name"`
	CountTeam   int16  `json:"countTeam" gorm:"column:count_team"`
	PicName     string `json:"picName" gorm:"column:pic_name"`
	PicEmail    string `json:"picEmail" gorm:"column:pic_email"`
	CountBugs   int16  `json:"countBugs" gorm:"column:count_bugs"`
	CountModule int32  `json:"countModule" gorm:"column:count_module"`
}
