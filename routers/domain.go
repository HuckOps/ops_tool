package routers

import (
	"ops_tool/controller/domain"

	"github.com/gin-gonic/gin"
)

func DomainRoute(e *gin.Engine) {
	DomainMap := e.Group("/api/domain")
	{
		DomainMap.GET("/whois/:domain", domain.GetWhoisInfo)
		DomainMap.GET("/icp/:domain", domain.ICPInfo)
	}
}
