package routers

import (
	"ops_tool/controller/domain"
	"ops_tool/controller/ws"

	"github.com/gin-gonic/gin"
)

func WebSocket(e *gin.Engine) {
	wsGroup := e.Group("/api/ws")
	{
		wsGroup.GET("/dns/:domain", ws.WebsocketManager.WsClient, domain.DNSResolve)
	}
}
