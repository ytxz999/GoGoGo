package controller

import (
	"memo/common"
	"memo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// CreateTodo 事项创建
//
// @Summary 创建todo
// @Description 用户输入title,content,start_time,end_time实现创建事项
// @Tags Todo
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body CreateTodoRequest true "todo信息"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/todos [post]
func CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		common.BadReqFail(c, "参数错误")
		return
	}

	value, exist := c.Get("userId")
	if !exist {
		common.UnauFail(c, "用户不存在")
	}
	userId := value.(uint)

	err := service.CreateTodo(req.Title, req.Content, req.StartTime, req.EndTime, userId)
	if err != nil {
		common.InternalFail(c, "创建todo失败")
		return
	}

	// 返回注册结果
	common.Success(c, "创建成功")
}

type QueryTodoRequest struct {
	Page    *int   `json:"page"`
	Status  *int   `json:"status"`
	Keyword string `json:"keyword"`
}

// GetTodoList 事项查询
//
// @Summary 查询todo
// @Description 用户输入status,key实现查询事项
// @Tags Todo
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body QueryTodoRequest true "todo信息"
// @Success 200 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/todos [get]
func GetTodoList(c *gin.Context) {
	value, exist := c.Get("userId")
	if !exist {
		common.UnauFail(c, "用户不存在")
		return
	}
	var req QueryTodoRequest
	if page := c.Query("page"); page != "" {
		s, _ := strconv.Atoi(page)
		req.Page = &s
	}
	size := 10
	//查询status是否为空
	if status := c.Query("status"); status != "" {
		s, _ := strconv.Atoi(status)
		req.Status = &s
	}
	req.Keyword = c.DefaultQuery("keyword", "")

	userId := value.(uint)
	todoList, total, err := service.GetTodoList(req.Page, size, req.Status, req.Keyword, userId)
	if err != nil {
		common.InternalFail(c, "获取todo失败")
		return
	}
	common.Success(c, gin.H{
		"data":   todoList,
		"total":  total,
		"page":   req.Page,
		"size":   size,
		"msg":    "ok",
		"status": 200,
	})
}

type UpdateTodoRequest struct {
	Id     *uint `json:"id"`
	Status int   `json:"status"`
}

// UpdateTodo 事项更新
//
// @Summary 更新todo
// @Description 用户输入id,status实现更新事项
// @Tags Todo
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body UpdateTodoRequest  true "todo信息"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/todos [put]
func UpdateTodo(c *gin.Context) {
	value, exist := c.Get("userId")
	if !exist {
		common.UnauFail(c, "用户不存在")
	}
	userId := value.(uint)
	var req UpdateTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		common.BadReqFail(c, "参数错误")
	}
	err := service.UpdateTodo(req.Id, req.Status, userId)
	if err != nil {
		common.InternalFail(c, "更新todo失败")
		return
	}
	common.Success(c, "更新成功")

}

type DeleteTodoRequest struct {
	Id     *uint `json:"id"`
	Status *int  `json:"status"`
	//确认参数，保证全部删除为正确
	All bool `json:"all"`
}

// DeleteTodo 事项删除
//
// @Summary 删除todo
// @Description 用户输入id,status，all实现删除事项
// @Tags Todo
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body DeleteTodoRequest  true "todo信息"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/todos [delete]
func DeleteTodo(c *gin.Context) {
	value, exist := c.Get("userId")
	if !exist {
		common.UnauFail(c, "用户不存在")
	}
	userId := value.(uint)
	var req DeleteTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		common.BadReqFail(c, "参数错误")
	}
	err := service.DeleteTodo(req.Id, req.Status, req.All, userId)
	if err != nil {
		common.InternalFail(c, "删除todo失败")
		return
	}
	common.Success(c, "删除成功")
}
