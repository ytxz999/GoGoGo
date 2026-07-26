package controller

// 表示层
// 职责：负责 HTTP 请求接收、请求参数校验与绑定、响应数据序列化与状态码返回。
import (
	"memo/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 实现注册功能
func Register(c *gin.Context) {
	// 解析请求参数
	var req RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}

	// 调用服务层注册功能
	err := service.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "用户注册失败",
			"err": err.Error(),
		})
		return

	}

	// 返回注册结果
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 实现登录功能
func Login(c *gin.Context) {
	var req LoginRequest
	// 解析请求参数
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}
	// 调用服务层登录功能
	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "登录失败",
			"err": err.Error(),
		})
		return
	}

	// 返回登录结果
	c.JSON(200, gin.H{
		"msg":   "登录成功",
		"token": token,
	})

}
