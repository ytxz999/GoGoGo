package controller

// 表示层
// 职责：负责 HTTP 请求接收、请求参数校验与绑定、响应数据序列化与状态码返回。
import (
	"memo/common"
	"memo/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register 用户注册
//
// @Summary 注册User
// @Description 用户输入用户名和密码实现注册
// @Tags User
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "注册信息"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/register [post]
func Register(c *gin.Context) {
	// 解析请求参数
	var req RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		common.BadReqFail(c, "参数错误")
		return
	}

	// 调用服务层注册功能
	err := service.Register(req.Username, req.Password)
	if err != nil {
		common.InternalFail(c, "用户注册失败")
		return

	}

	// 返回注册结果
	common.Success(c, "注册成功")

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 用户登录
//
// @Summary 登录User
// @Description 用户输入用户名和密码获取JWT Token
// @Tags User
// @Accept json
// @Produce json
// @Param request body LoginRequest true "登录信息"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/login [post]
// 实现登录功能
func Login(c *gin.Context) {
	var req LoginRequest
	// 解析请求参数
	if err := c.ShouldBind(&req); err != nil {
		common.BadReqFail(c, "参数错误")
		return
	}
	// 调用服务层登录功能
	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		common.InternalFail(c, "登录失败")
		return
	}

	// 返回登录结果
	common.Success(c, token)
}
