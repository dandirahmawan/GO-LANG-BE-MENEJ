package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string
	Title  string
	Author string
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust"},
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust"},
}

func GetBooks(ctx *gin.Context) {
	b := books[0]
	ctx.IndentedJSON(http.StatusOK, b)
}
