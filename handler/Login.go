package handler

import (
	"net/http"

	"github.com/dandirahmawan/menej_api_go/services"
	"github.com/gin-gonic/gin"
)

type formLogin struct {
	Email    string `form:"email" json:"user"`
	Password string `form:"password" json:"password"`
}

func Login(ctx *gin.Context) {
	var form formLogin
	ctx.Bind(&form)

	email := form.Email
	password := form.Password
	data := services.FindByEmailAndPassword(email, password)
	ctx.IndentedJSON(http.StatusOK, data)
}
