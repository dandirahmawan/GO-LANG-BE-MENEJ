package commons

import (
	"math/rand"
	"strings"
	"time"

	"github.com/dandirahmawan/menej_api_go/config"
	"github.com/dandirahmawan/menej_api_go/model"
)

const charset = "abcdefghuvwx0987yzijklmnopQRSTU231ABqrs546tCDEFGHIbcDBcdJKLMNOPQRSTUVWXYZ"

func GeneratdUUID(n int) string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func GetDataSession(sessionid string) model.SessionModel {
	DB, _ := config.ConnectDB()
	type Model model.SessionModel
	var data model.SessionModel
	DB.Where(Model{Id: sessionid}).Find(&data)

	return data
}
