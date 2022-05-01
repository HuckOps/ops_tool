package main

import (
	"ops_tool/config"
	"ops_tool/controller/ws"
	"ops_tool/db"
	"ops_tool/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendAllService()
	go ws.WebsocketManager.SendAllService()

	config.InitConfig()
	db.InitMySQL()
	r := gin.Default()
	routers.Register(r)
	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}
