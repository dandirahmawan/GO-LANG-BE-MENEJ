package commons

import (
	"math/rand"
	"strings"
	"time"

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
	type Model model.SessionModel
	var data model.SessionModel
	model.DB.Where(Model{Id: sessionid}).Find(&data)

	// defer func() {
	// 	config.CloseConnection(DB)
	// 	config.CloseConnection(exc)
	// 	fmt.Println("close connection get data session")
	// }()

	return data
}
