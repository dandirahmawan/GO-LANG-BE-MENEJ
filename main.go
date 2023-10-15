package main

import (
	"github.com/dandirahmawan/menej_api_go/model"
	router "github.com/dandirahmawan/menej_api_go/routers"
)

func main() {
	db := model.ConnectDB()
	routerInit := router.InitRouter(db)
	routerInit.Run("localhost:8888")
}
