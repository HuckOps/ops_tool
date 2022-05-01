package restapi

import "github.com/gin-gonic/gin"

type StatusCode int

// 状态码定义

const (
	OK           StatusCode = 200
	BadRequest   StatusCode = 400
	UnAuthorized StatusCode = 403
	NotFound     StatusCode = 404
	ServerError  StatusCode = 500
)

type Code int

// 响应状态

const (
	Success Code = iota
	Failed
)

// 响应结构

type Response struct {
	Code Code        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (response *Response) Response(c *gin.Context, statusCode StatusCode) {
	c.JSON(int(statusCode), response)
}
