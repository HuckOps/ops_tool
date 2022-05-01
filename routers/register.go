package routers

import "github.com/gin-gonic/gin"

func Register(e *gin.Engine) {
	WebSocket(e)
	IPPoolRoute(e)
	DomainRoute(e)
}
