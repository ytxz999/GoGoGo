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

// 创建todo
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
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Status  *int   `json:"status"`
	Keyword string `json:"keyword"`
}

// 获取todo列表
func GetTodoList(c *gin.Context) {
	value, exist := c.Get("userId")
	if !exist {
		common.UnauFail(c, "用户不存在")
		return
	}
	var req QueryTodoRequest
	req.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	req.Size, _ = strconv.Atoi(c.DefaultQuery("size", "10"))
	//查询status是否为空
	if status := c.Query("status"); status != "" {
		s, _ := strconv.Atoi(status)
		req.Status = &s
	}
	req.Keyword = c.DefaultQuery("keyword", "")

	userId := value.(uint)
	todoList, total, err := service.GetTodoList(req.Page, req.Size, req.Status, req.Keyword, userId)
	if err != nil {
		common.InternalFail(c, "获取todo失败")
		return
	}
	common.Success(c, gin.H{
		"data":   todoList,
		"total":  total,
		"page":   req.Page,
		"size":   req.Size,
		"msg":    "ok",
		"status": 200,
	})
}

type UpdateTodoRequest struct {
	Id     *uint `json:"id"`
	Status int   `json:"status"`
}

// 更新todo
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

// 删除todo
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
