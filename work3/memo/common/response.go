package common

import "github.com/gin-gonic/gin"

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Msg:  "success",
		Data: data,
	})
}

func BadReqFail(c *gin.Context, data interface{}) {
	c.JSON(400, Response{
		Msg:  "fail",
		Data: data,
	})
}

func UnauFail(c *gin.Context, data interface{}) {
	c.JSON(401, Response{
		Msg:  "fail",
		Data: data,
	})
}

func NotFoundFail(c *gin.Context, data interface{}) {
	c.JSON(404, Response{
		Msg:  "fail",
		Data: data,
	})
}

func InternalFail(c *gin.Context, data interface{}) {
	c.JSON(500, Response{
		Msg:  "fail",
		Data: data,
	})
}
