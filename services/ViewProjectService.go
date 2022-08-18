package services

import (
	"github.com/dandirahmawan/menej_api_go/config"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

var DB, _ = config.ConnectDB()

func GetAllViewProjects(ctx *gin.Context) []model.ViewProjectModel {
	userid := ctx.Request.Header.Get("userid")
	where := " AND (p.pic = '" + userid + "' OR p.project_id IN (SELECT \r\n" +
		"											pt.project_id \r\n" +
		"										FROM project_team pt \r\n" +
		"										WHERE pt.user_id = '" + userid + "' AND (pt.user_status IS NULL OR pt.user_status != 'N'))) \r\n" +
		"ORDER BY p.created_date ASC"

	rawQuery := model.RawQueryViewProject + where
	// fmt.Println(rawQuery)
	var data []model.ViewProjectModel
	DB.Raw(rawQuery).Scan(&data)
	return data
}
