package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dandirahmawan/menej_api_go/constanta"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var UrlNoAauth [1]
// UrlNoAauth =  = [1]{"/login"}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/*get data heaeder*/
		userid := c.Request.Header.Get("userid")
		sessionid := c.Request.Header.Get("sessionid")

		url := c.Request.URL.String()
		var isNonAuth bool = isUrlNonAuth(url)

		/*jika url yang diakses adalah url yang tidak membutuhkan
		authorization maka akan melewati fase ini*/
		if !isNonAuth {
			if userid == "" || sessionid == "" {
				c.IndentedJSON(http.StatusForbidden, "Require data header userid, sessionid")
				c.Abort()
				return
			}

			/*validate token or sessionid*/
			sessionid := c.Request.Header.Get("sessionid")
			var m model.SessionModel
			m.Id = sessionid
			data := m.FindBySessionid()
			if len(data) == 0 {
				c.IndentedJSON(http.StatusUnauthorized, "Session is expired")
				c.Abort()
				return
			}
			// fmt.Println("sessionid is ", sessionid)
		}

		c.Next()
	}
}

func isUrlNonAuth(url string) bool {
	var result bool = false

	for i := 0; i < len(constanta.UrlNotAuth); i++ {
		item := constanta.UrlNotAuth[i]
		fmt.Println(item, url)
		if strings.Contains(url, item) {
			result = true
			break
		}
	}

	return result
}
