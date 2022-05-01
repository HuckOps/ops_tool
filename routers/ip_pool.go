package routers

import (
	ippool "ops_tool/controller/ip_pool"

	"github.com/gin-gonic/gin"
)

func IPPoolRoute(e *gin.Engine) {
	IPPoolMap := e.Group("/api/ip_tool")
	{
		IPPoolMap.GET("/", ippool.GetIPInfoByIP)
		IPPoolMap.GET("/:ip", ippool.GetIPInfoByIP)
	}
}
