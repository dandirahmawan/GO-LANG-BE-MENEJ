package middleware

import (
	"net/http"
	"strings"

	"github.com/dandirahmawan/menej_api_go/config"
	"github.com/dandirahmawan/menej_api_go/constanta"
	"github.com/gin-gonic/gin"
)

// var UrlNoAauth [1]
// UrlNoAauth =  = [1]{"/login"}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
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
		}

		_, err := config.ConnectDB()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			panic("Database connection is failure")
		}

		c.Next()
	}
}

func isUrlNonAuth(url string) bool {
	var result bool = false

	for i := 0; i < len(constanta.UrlNotAuth); i++ {
		item := constanta.UrlNotAuth[i]

		if strings.Contains(url, item) {
			result = true
			break
		}
	}

	return result
}
